import http from 'k6/http';
import { check } from 'k6';

export const options = {
  stages: [
    { duration: '1m', target: 100 },  
    { duration: '2m', target: 500 },  
    { duration: '1m', target: 1000 }, 
    { duration: '1m', target: 5000 }, 
    { duration: '2m', target: 1000 }, 
    { duration: '1m', target: 500 },  
    { duration: '1m', target: 100 },  
  ],
  thresholds: {
    http_req_duration: ['p(95)<500'],
  },
};

export default function () {
  const response = http.get('https://shortener.solucionesvaltech.com/VsgjlHZNR');
  
  check(response, {
    'is status 200': (r) => r.status === 200,
  });
}
