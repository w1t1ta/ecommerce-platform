import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import './index.css'; // นำเข้าไฟล์ index.css เพื่อใช้สไตล์หลักของแอป
import 'bootstrap/dist/css/bootstrap.min.css'; // นำเข้า Bootstrap เพื่อใช้สไตล์พื้นฐาน
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

reportWebVitals();
