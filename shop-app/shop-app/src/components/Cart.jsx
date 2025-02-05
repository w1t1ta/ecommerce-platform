import React, { useState } from "react";

import Header from './Header';
import Footer from './Footer';
import './css/Cart.css'

const Cart = () => {
  const [cartItems, setCartItems] = useState([
    { id: 1, name: 'Practical DEVOPS and Cloud Engineering book', price: 545, quantity: 1 },
    { id: 2, name: 'Canson สมุดวาดเขียนสีน้ำ A5 300 แกรม ชนิดเรียบ', price: 106, quantity: 1 },
  ]);
  const [coupon, setCoupon] = useState('');
  const shippingCost = 0; // ค่าจัดส่งฟรี

  const handleQuantityChange = (id, quantity) => {
    const updatedItems = cartItems.map((item) =>
      item.id === id ? { ...item, quantity: Number(quantity) } : item
    );
    setCartItems(updatedItems);
  };

  const calculateSubtotal = () => {
    return cartItems.reduce((total, item) => total + item.price * item.quantity, 0);
  };

  // เพิ่มฟังก์ชัน calculateTotalItems ที่นี่
  const calculateTotalItems = () => {
    return cartItems.reduce((total, item) => total + item.quantity, 0);
  };

  const subtotal = calculateSubtotal();
  const total = subtotal + shippingCost;
  const totalItems = calculateTotalItems();

  return (
    <div className="page-container">
      <Header /> {/* เพิ่ม Header ที่นี่ */}

      <main className="cart-content">
        <h2>Your Cart</h2>
        <div className="cart-items-list">
          {cartItems.map((item, index) => (
            <div key={item.id} className="cart-item-row">
              <div>
                {/* ใช้ item.id สำหรับแสดงข้อมูล ID ของสินค้า */}
                {item.id === 1 && (
                  <div>
                    
                    <img src="/image/devops.png" width="45" height="45" alt="DEVOPS Book" />
                  </div>
                )}

                {item.id === 2 && (
                  <div>
                    
                    <img src="/image/canson.png" width="45" height="45" alt="Canson Book" />
                  </div>
                )}
              </div>

              <span>{index + 1}. {item.name}</span>
              <span>${item.price}</span>
              <input
                type="number"
                value={item.quantity}
                onChange={(e) => handleQuantityChange(item.id, e.target.value)}
                min="1"
              />
              <span>${item.price * item.quantity}</span>
            </div>
          ))}
        </div>

        {/* Container สำหรับ Cart Total */}
        <div className="cart-total-container">
          <h3>Cart Total</h3>
          <div className="cart-total-details">
            <div className="subtotal">
              <span>Subtotal:</span>
              <span>${subtotal}</span>
            </div>
            <div className="shipping">
              <span>Shipping:</span>
              <span>Free</span>
            </div>
            <div className="total">
              <span>Total:</span>
              <span>${total}</span>
            </div>
          </div>

          {/* ปุ่ม Process to Checkout */}
          <button className="checkout-button">Process to Checkout</button>
        </div>
      </main>

      <Footer /> {/* เพิ่ม Footer ที่นี่ */}
    </div>
  );
};

export default Cart;
