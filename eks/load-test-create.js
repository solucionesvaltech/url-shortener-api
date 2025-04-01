import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
  stages: [
    { duration: '30s', target: 500 }, 
    { duration: '30s', target: 1000 },
    { duration: '1m', target: 1000 }, 
    { duration: '15s', target: 1500 }, 
    { duration: '10s', target: 5000 }, 
    { duration: '30s', target: 0 },   
  ],
};

export default function () {
  const url = 'https://shortener.solucionesvaltech.com/urls';
  const payload = JSON.stringify({ url: 'https://example.com' });
  const params = { headers: { 'Content-Type': 'application/json' } };

  let res = http.post(url, payload, params);

  check(res, {
    'is status 201': (r) => r.status === 201,
  });

  sleep(1);
}
