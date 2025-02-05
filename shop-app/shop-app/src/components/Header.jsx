import React, { useEffect, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { FaSearch, FaShoppingCart, FaUser } from 'react-icons/fa';
import { Button, Modal, Dropdown } from 'react-bootstrap'; // เพิ่ม Dropdown จาก react-bootstrap
import axios from 'axios';
import './css/Header.css';

const Header = () => {
    const [cartItems, setCartItems] = useState(0); // จำนวนสินค้าในตะกร้า
    const [showLogin, setShowLogin] = useState(false);
    const [userName, setUserName] = useState(''); // ชื่อผู้ใช้
    const navigate = useNavigate();

    useEffect(() => {
        // ดึงข้อมูลตะกร้าสินค้า
        axios.get('')
            .then(response => {
                setCartItems(response.data.totalItems);
            })
            .catch(error => {
                console.error("Error fetching cart items:", error);
            });

        // ดึงข้อมูลชื่อผู้ใช้
        axios.get('')
            .then(response => {
                setUserName(response.data.name);
            })
            .catch(error => {
                console.error("Error fetching user data:", error);
            });
    }, []);

    // ฟังก์ชันสำหรับล็อกเอาท์
    const handleLogout = () => {
        setUserName('');
        navigate('/login');
    };

    return (
        <header className="header">
       
       <div className="logo">
    <Link to="/" className="logo-link">MAGE</Link>
</div>
           
            <nav>
                <ul>
                    <li><Link to="/">Home</Link></li>
                    <li><Link to="/contact">Contact</Link></li>
                    <li><Link to="/aboutpage">About</Link></li>
                    <li><Link to="/product">Shop</Link></li>
                    
                    
                </ul>
            </nav>
            <div className="search">
                <input type="text" placeholder="What are you looking for?" />
                <button className="search-button">
                    <Link to="/search">
                        <img src="https://cdn-icons-png.flaticon.com/128/54/54481.png" style={{ width: '12px', height: '12px' }} alt="Search Icon"/>
                    </Link>
                </button>
            </div>
            <div className="icons">
                <span>
                    <Link to="/cart">
                        <img src="https://i.pinimg.com/originals/7f/24/92/7f249252404646c08d90976505cb6937.jpg" style={{ width: '30px', height: '30px' }} alt="Cart Icon"/>
                     
                    </Link>
                </span>
                <span className="user-profile">
    {userName ? (
        <span className="user-name">
            <img 
                src="https://i.pinimg.com/564x/02/72/3a/02723a8b181c646ad15095dd4786dac1.jpg" 
                style={{ width: '22px', height: '23px', marginRight: '5px' }} 
                alt="User Icon"
            />
            Hello, {userName}
            <button onClick={handleLogout} className="logout-button">Logout</button>
        </span>
    ) : (
        <Link to="/login" className="login-link">
            <img 
                src="https://i.pinimg.com/564x/02/72/3a/02723a8b181c646ad15095dd4786dac1.jpg" 
                style={{ width: '22px', height: '23px', marginRight: '5px' }} 
                alt="User Icon"
            />
           
        </Link>
    )}
</span>
            </div>
        </header>
    );
};

export default Header;
