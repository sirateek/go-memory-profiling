import http from "k6/http";
import { check } from "k6";
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export const options = {
  scenarios: {
    open_model: {
      executor: "constant-arrival-rate",
      rate: 1000,
      timeUnit: "1s",
      duration: "1m",
      preAllocatedVUs: 20,
    },
  },
  noConnectionReuse: true,
  userAgent: "K6LoadTest/1.0",
};

export default function () {
  const reqID = uuidv4();
  let res = http.post(
    "http://localhost:30999/example",
    {},
    {
      headers: {
        "X-REQ-ID": reqID,
      },
    }
  );

  check(res, {
    "is status 200": (r) => r.status === 200,
  });

  check(res, {
    "is response correctly": (r) => r.body === reqID,
  });
}
