import { Request, Response, NextFunction } from 'express';

const checkAuth = (req: Request, res: Response, next: NextFunction) => {
  if (!req.isAuthenticated()) {
    return res.status(401).json({ error: 'You must log in!' });
  }
  next();
};

export default checkAuth;
