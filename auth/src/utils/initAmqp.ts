import { HttpError } from '@adwesh/common';
import amqplib from 'amqplib/callback_api';

const queue = 'auth_topic';

let _conn: amqplib.Connection;
export class AmqpClient {
  get client(): amqplib.Connection {
    if (!_conn) {
      throw new HttpError('AMQP must be initialised', 500);
    }
    return _conn;
  }

  connect(amqpUri: string) {
    amqplib.connect(amqpUri, (err, connection) => {
      if (err)
        throw new HttpError(
          err instanceof HttpError ? err.message : 'error connecting to AMQP',
          500
        );
      _conn = connection;
      console.log('Connected to AMQP . . .');
    });
  }

  sendToQueue(data: { userId: string; expiry: number }) {
    if (_conn) {
      _conn.createChannel((err, channel) => {
        if (err) throw err;
        channel.assertQueue(queue);
        channel.sendToQueue(queue, Buffer.from(JSON.stringify(data)));
      });
    } else {
      throw new HttpError('Error initialising connection', 500);
    }
  }
}
