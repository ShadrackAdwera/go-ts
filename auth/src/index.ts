import mongoose from 'mongoose';
import { app } from './app';
import { AmqpClient } from './utils/initAmqp';

if (!process.env.COOKIE_KEY) {
  throw new Error('COOKIE_KEY is not defined!');
}

if (!process.env.MONGO_URI) {
  throw new Error('MONGO URI is not defined!');
}

if (!process.env.REDIS_HOST) {
  throw new Error('REDIS HOST is not defined!');
}

if (!process.env.GOTS_CLIENT_ID) {
  throw new Error('GOTS_CLIENT_ID must be defined');
}
if (!process.env.GOTS_CLIENT_SECRET) {
  throw new Error('GOTS_CLIENT_SECRET must be defined');
}

if (!process.env.AMQP_CONNECTION_STRING) {
  throw new Error('GOTS_CLIENT_SECRET must be defined');
}

const startApp = async () => {
  try {
    await mongoose.connect(process.env.MONGO_URI!, {
      auth: {
        username: 'admin',
        password: 'password',
      },
      dbName: 'auth',
    });
    await new AmqpClient().connect(process.env.AMQP_CONNECTION_STRING!);
    app.listen(5000);
    console.log('Auth Service Started on PORT: 5000');
  } catch (error) {
    console.log(error instanceof Error ? error.message : error);
    throw error;
  }
};

startApp();
