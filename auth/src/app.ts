import express, { Request, Response, NextFunction } from 'express';
import passport from 'passport';
import session, { SessionOptions } from 'express-session';
import MongoStore from 'connect-mongo';

import { authRoutes } from './routes/user-routes';
import { HttpError } from '@adwesh/common';
import { COOKIE_MAX_AGE } from './utils/constants';
import './utils/Passport';

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
    dbName: 'auth',
    mongoOptions: {
      auth: {
        username: 'admin',
        password: 'password',
      },
    },
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

app.use((_req: Request, _res: Response, _next: NextFunction) => {
  throw new HttpError('This method / route does not exist!', 404);
});

app.use(
  (error: HttpError, _req: Request, res: Response, next: NextFunction) => {
    if (res.headersSent) {
      return next(error);
    }
    res
      .status(error.code || 500)
      .json({ message: error.message || 'An error occured, try again!' });
  }
);

export { app };
