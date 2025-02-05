import React from 'react';
import { useNavigate } from 'react-router-dom';  // นำเข้า useNavigate

import './css/ProductDetail.css';

const ProductDetailsPage = () => {
    const navigate = useNavigate();  // สร้างตัวแปร navigate
  
    // ฟังก์ชันสำหรับจัดการการคลิกปุ่ม "View Details"
    const handleViewDetails = (productId) => {
      console.log("คลิกปุ่ม View Details แต่ไม่ทำการนำทาง");
    };
  const product = {
    name: "บุกทะลวงข้อสอบTOEIC",
    price: "฿750",
    description: "ฝึกให้ชัวร์ด้วยข้อสอบ TOEIC เสมือนจริงเหมาะทั้งเรียนรู้ด้วยตัวเอง หรือติวเตอร์ใช้ประกอบการสอนข้อสอบ TOEIC รูปแบบใหม่ เสมือนจริง 5 ชุด 1000 ข้อ จุใจ รวมข้อสอบทั้งพาร์ Listening + Readingเหมาะสำหรับผู้สอบที่มุ่งเป๊าคะแนน 500+ถึง 800+",
    colorOptions: ["Black", "Red", "Blue"],
    sizeOptions: ["Standard", "Large"],
    stock: 10,
    relatedProducts: [
      { id: 1, name: "เถียงอย่างไรให้ชนะแมว", price: "฿175", image: "book4.JPG"},
      { id: 2, name: "ร้านขายยาปริศนารับแก้ปัญหาหัวใจ", price: "฿385", image: "book5.JPG"},
      { id: 3, name: "Practical Devops and Cloud Engineering", price: "฿545", image: "book3.JPG"},
    ],
    image: "book.JPG"//ไม่ได้สร้างไฟล์เก็บรูปนะ เราเอารูปลง
  };

  return (
    <div className="w-full p-4 h-screen overflow-y-auto">
      <div className="product-details-page" style={{ maxWidth: '1440px', margin: 'auto' }}>
        <div className="breadcrumb">
          <p> </p>
          <p> </p>
          <a href="/">Account</a> / <a href="/">Book</a> / {product.name}
          <p> </p>
          <p> </p>
          <p> </p>
        </div>
        <div className="product-details" style={{ display: 'flex', justifyContent: 'center', alignItems: 'center', flexWrap: 'wrap' }}>
          {/* Product Image */}
          <div className="product-image" style={{ flex: '1 1 auto', display: 'flex', justifyContent: 'center' }}>
            <img
              src={product.image}
              alt={product.name}
              style={{ maxWidth: '100%', width: '400px', height: 'auto', margin: 'auto' }} 
            />
            <p> </p>
            <p> </p>
          </div>
          {/* Product Info (on the right) */}
          <div className="product-info" style={{ flex: '1 1 auto', maxWidth: '300px', textAlign: 'left', paddingLeft: '20px' }}> 
            <h1>{product.name}</h1>
            <p className="product-price">{product.price}</p>
            <p className="product-description">{product.description}</p>
          
            <div className="product-options">
              <div className="product-color">
                <h3>Color</h3>
                <select defaultValue="Black">
                  {product.colorOptions.map((color, index) => (
                    <option key={index} value={color}>{color}</option>
                  ))}
                </select>
              </div>

              <div className="product-size">
                <h3>Size</h3>
                <select defaultValue="Standard">
                  {product.sizeOptions.map((size, index) => (
                    <option key={index} value={size}>{size}</option>
                  ))}
                </select>
              </div>

              <div className="product-quantity">
                <h3>Quantity</h3>
                <input type="number" min="1" max={product.stock} defaultValue="1" />
              </div>
            </div>

            <button className="buy-now">Buy Now</button>
          </div>
        </div>

        <div className="related-products">
          <h2>Related Products</h2>
          {product.relatedProducts.map(related => (
            <div key={related.id} className="related-product-item" style={{ display: 'inline-block', width: 'calc(33.333% - 20px)', margin: '10px' }}>
              <img src={related.image} alt={related.name} style={{width: '100%', height: 'auto', maxWidth: '100%', width: '200px', height: 'auto', margin: '0px'}} />
              <p>{related.name}</p>
              <p>{related.price}</p>

              {/* ปุ่ม View Details */}
              <button className="view-details-btn" onClick={() => handleViewDetails(related.id)}>View Details</button>

            </div>
          ))}
        </div>
      </div>
    </div>
  );
};
export default ProductDetailsPage;
