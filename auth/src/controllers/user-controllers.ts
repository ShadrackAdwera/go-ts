import { HttpError } from '@adwesh/common';
import { Request, Response, NextFunction } from 'express';
import { Types } from 'mongoose';
import { User, UserDoc } from '../models/User';

const getAuthCallback = (req: Request, res: Response, next: NextFunction) => {
  return res.redirect('/auth/users');
};

const logOut = (req: Request, res: Response, next: NextFunction) => {
  req.logOut({ keepSessionInfo: false }, (err) => {
    if (err) {
      return next(
        new HttpError(
          err instanceof HttpError ? err.message : 'Error logging you out',
          500
        )
      );
    }
  });
  res.redirect('/login');
};

const currentUser = (req: Request, res: Response, next: NextFunction) => {
  return res.status(200).json({ user: req.user });
};

const fetchUsers = async (req: Request, res: Response, next: NextFunction) => {
  let users: (UserDoc & { _id: Types.ObjectId })[];

  try {
    users = await User.find();
  } catch (error) {
    return next(
      new HttpError(
        error instanceof Error ? error.message : 'An error occured : 29',
        500
      )
    );
  }

  res.status(200).json({ count: users.length, users });
};

export { getAuthCallback, logOut, fetchUsers, currentUser };
