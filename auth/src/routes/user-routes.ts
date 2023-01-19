import { checkAuth } from '@adwesh/common';
import { Router } from 'express';
import passport from 'passport';

import {
  currentUser,
  fetchUsers,
  getAuthCallback,
  logOut,
} from './../controllers/user-controllers';

const router = Router();

router.get(
  '/google/callback',
  passport.authenticate('google'),
  getAuthCallback
);
router.use(checkAuth);
router.get('/current-user', currentUser);
router.get('/users', fetchUsers);
router.get('/logout', logOut);

export { router as authRoutes };
