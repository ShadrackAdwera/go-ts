import { HttpError } from '@adwesh/common';
import express, { Request, Response, NextFunction } from 'express';

import { authRoutes } from './routes/user-routes';

const app = express();

app.use(express.json());
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
