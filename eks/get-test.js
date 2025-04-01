import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 1000 },  // Escalamos hasta 1000 VUs
    { duration: '2m', target: 1000 },
    { duration: '1m', target: 5000 },  // Aumentamos hasta 5000 VUs
    { duration: '2m', target: 5000 },
    { duration: '1m', target: 0 },     // Disminuimos gradualmente
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],  // El 95% de las solicitudes deben responder en menos de 500 ms
    http_req_failed: ['rate<0.01'],    // Menos del 1% de fallos permitidos
  },
};

const BASE_URL = 'https://shortener.solucionesvaltech.com';

export default function () {
  const shortID = 'JOr5DFWNR';
  const res = http.get(`${BASE_URL}/${shortID}`);
  
  check(res, {
    'is status 200': (r) => r.status === 200,
  });
  
  sleep(1);
}
