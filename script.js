import http from 'k6/http';
import { check } from 'k6';

export const options = {
  vus: 5,
  duration: '10s',
};

export default function () {
  const res = http.get('https://httpbin.test.k6.io/get');
  check(res, {
    'status is 200': (r) => r.status === 200,
  });
}