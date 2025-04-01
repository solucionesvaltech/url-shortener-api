import http from 'k6/http';
import { check } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 100 },
    { duration: '2m', target: 500 },
    { duration: '1m', target: 1000 },
    { duration: '3m', target: 4500 },
    { duration: '2m', target: 1000 },
    { duration: '1m', target: 500 },
    { duration: '30s', target: 100 },
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],
  },
};

export default function () {
  const response = http.get('https://shortener.solucionesvaltech.com/dLu13HWHR');
  
  check(response, {
    'is status 200': (r) => r.status === 200,
  });
}