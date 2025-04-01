import http from 'k6/http';
import { check } from 'k6';

export const options = {
  scenarios: {
    steady_load: {
      executor: 'constant-arrival-rate',
      rate: 83,
      timeUnit: '1s',
      duration: '5m',
      preAllocatedVUs: 100,
      maxVUs: 500,
    },
  },
};

export default function () {
  const response = http.get('https://shortener.solucionesvaltech.com/VsgjlHZNR');
  
  check(response, {
    'is status 200': (r) => r.status === 200,
  });
}
