import { Document, Schema, model } from 'mongoose';

export interface UserDoc extends Document {
  username: string;
  email: string;
  googleId: string;
}

const userSchema = new Schema<UserDoc>(
  {
    username: { type: String, required: true },
    email: { type: String, required: true },
    googleId: { type: String, required: true },
  },
  { timestamps: true, toJSON: { getters: true } }
);

const User = model<UserDoc>('user', userSchema);

export { User };
