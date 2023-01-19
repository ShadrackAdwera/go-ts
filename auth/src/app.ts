import { HttpError } from '@adwesh/common';
import express, { Request, Response, NextFunction } from 'express';
import passport from 'passport';
import session, { SessionOptions } from 'express-session';
import MongoStore from 'connect-mongo';

import { authRoutes } from './routes/user-routes';
import { COOKIE_MAX_AGE } from './utils/constants';

const app = express();

app.use(express.json());

const sess: SessionOptions = {
  secret: process.env.COOKIE_KEY!,
  resave: false,
  saveUninitialized: false,
  cookie: {
    secure: app.get('env') === 'production',
    maxAge: COOKIE_MAX_AGE,
  },
  store: new MongoStore({
    collectionName: 'session',
    mongoUrl: process.env.MONGO_URI,
  }),
};

if (app.get('env') === 'production') {
  app.set('trust proxy', 1); // trust first proxy
  sess.cookie!.secure = true; // serve secure cookies
}

app.use(session(sess));
app.use(passport.initialize());
app.use(passport.session());
app.use('/auth', authRoutes);

app.use((req: Request, res: Response, next: NextFunction) => {
  throw new HttpError('Invalid Method / Route', 404);
});

app.use((error: Error, req: Request, res: Response, next: NextFunction) => {
  if (res.headersSent) return next(error);
  res.status(error instanceof HttpError ? error.code : 500).json({
    message:
      error instanceof HttpError ? error.message : 'Internal Server Error',
  });
});

export { app };
