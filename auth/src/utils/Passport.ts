import { Types } from 'mongoose';
import passport from 'passport';
import { Strategy as GoogleStrategy } from 'passport-google-oauth20';

import { User, UserDoc } from './../models/User';
import { BASE_URL } from './constants';

passport.serializeUser((user: any, done) => {
  done(undefined, user);
});

passport.deserializeUser((id, done) => {
  User.findById(id, (err: NativeError, user: UserDoc) => {
    done(undefined, user);
  });
});

passport.use(
  new GoogleStrategy(
    {
      callbackURL: `${BASE_URL}/auth/google/callback`,
      clientID: process.env.GOOGLE_CLIENT_ID!,
      clientSecret: process.env.GOOGLE_CLIENT_SECRET!,
      proxy: true,
    },
    async (accessToken, refreshToken, profile, done) => {
      let existingUser: (UserDoc & { _id: Types.ObjectId }) | null;
      try {
        existingUser = await User.findOne({ googleId: profile.id });
        if (existingUser) {
          // send to RabbitMQ
          return done(undefined, existingUser);
        }
        const user = await new User({
          googleId: profile.id,
          username: profile.displayName,
          email:
            profile.emails && profile.emails.length > 0
              ? profile.emails.toString()
              : '',
        }).save();
        // send to RabbitMQ
        done(null, user);
      } catch (err) {
        done(err instanceof Error ? err : 'An error occured', undefined);
      }
    }
  )
);
