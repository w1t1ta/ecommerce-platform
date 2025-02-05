


import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './css/Newproduct.css'; 


const ProductList = () => {
  const [products, setProducts] = useState([]);
  const [sellers, setSellers] = useState([]); // สร้าง state สำหรับเก็บข้อมูลผู้ขาย
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    
    const fetchProducts = axios.get('api/v1/products/new');
   
    const fetchSellers = axios.get('api/v1/sellers');

    
    axios
      .all([fetchProducts, fetchSellers])
      .then(
        axios.spread((productsResponse, sellersResponse) => {
          setProducts(productsResponse.data); 
          setSellers(sellersResponse.data);   
          setLoading(false); 
        })
      )
      .catch((err) => {
        setError(err.message); 
        setLoading(false);
      });
  }, []); 

  // ฟังก์ชันค้นหาชื่อผู้ขายโดยใช้ seller_id
  const getSellerName = (sellerId) => {
    const seller = sellers.find((seller) => seller.seller_id === sellerId);
    return seller ? seller.name : 'Unknown Seller';
  };

  if (loading) return <div className="loading">Loading...</div>; 
  if (error) return <div className="error">Error: {error}</div>; 

  return (
    <div className="product-list">
      <h1>สินค้าใหม่</h1>
      
      <div className="product-cards">
        {products.map((product) => (
          <div key={product.id} className="product-card">
            <img src={product.images[0]?.image_url} alt={product.name} className="product-image" />
            <h5>{product.name}</h5>
            <p className="product-price">ราคา : {product.price} บาท</p>
            {/* แสดงชื่อผู้ขายโดยเรียกฟังก์ชัน getSellerName */}
            <p className="product-seller">ร้านค้า : {getSellerName(product.seller_id)}</p>
          </div>
        ))}
      </div>
    </div>
  );
};

export default ProductList;
