import mongoose from 'mongoose';
import { app } from './app';

if (!process.env.COOKIE_KEY) {
  throw new Error('COOKIE_KEY is not defined!');
}

if (!process.env.MONGO_URI) {
  throw new Error('MONGO URI is not defined!');
}

if (!process.env.REDIS_HOST) {
  throw new Error('REDIS HOST is not defined!');
}

if (!process.env.GOOGLE_CLIENT_ID) {
  throw new Error('GOOGLE_CLIENT_ID must be defined');
}
if (!process.env.GOOGLE_CLIENT_SECRET) {
  throw new Error('GOOGLE_CLIENT_SECRET must be defined');
}

const startApp = async () => {
  try {
    await mongoose.connect(process.env.MONGO_URI!);
    app.listen(5000);
    console.log('Auth Service Started on PORT: 5000');
  } catch (error) {
    throw error;
  }
};

startApp();
