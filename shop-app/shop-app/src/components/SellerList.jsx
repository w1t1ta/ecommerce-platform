// src/components/SellerList.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './css/SellerList.css';

const SellerList = () => {
  const [sellers, setSellers] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // ดึงข้อมูลร้านค้าจาก API
    axios
      .get('api/v1/sellers')
      .then((response) => {
        setSellers(response.data); // ตั้งค่า sellers เป็นข้อมูลจาก API
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message); // ถ้ามีข้อผิดพลาด
        setLoading(false);
      });
  }, []);

  if (loading) return <div>กำลังโหลด...</div>; // แสดงข้อความกำลังโหลด
  if (error) return <div>ข้อผิดพลาด: {error}</div>; // ถ้ามีข้อผิดพลาด

  return (
    <div className="seller-container">
      <h1 className="seller-heading">ร้านค้าของเรา</h1>
      <div className="seller-list">
        {sellers.map((seller) => (
          <div key={seller.seller_id} className="seller-card">
            <h3 className="seller-name">{seller.name}</h3>
            <p className="seller-contact">Email : {seller.contact_info}</p>
            <Link to="/seller">ดูสินค้า</Link>
            
            
          </div>
        ))}
      </div>
    </div>
  );
};

export default SellerList;




