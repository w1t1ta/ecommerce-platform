// src/components/SignUp.js
import React, { useState } from 'react';
import axios from 'axios';
import './css/Register.css';

const Register = () => {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [message, setMessage] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            // ส่งข้อมูลไปยัง API สมัครสมาชิก
            await axios.post('/api/v1/user', { name, email, password });
            setMessage('Account created successfully');
        } catch (error) {
            setMessage('Error creating account');
        }
    };

    return (
        <div className="signup-container">
            <div className="signup-image">
                <img src="https://media.istockphoto.com/id/1414378934/th/%E0%B8%A3%E0%B8%B9%E0%B8%9B%E0%B8%96%E0%B9%88%E0%B8%B2%E0%B8%A2/%E0%B8%AD%E0%B8%B8%E0%B8%9B%E0%B8%81%E0%B8%A3%E0%B8%93%E0%B9%8C%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%99%E0%B8%AB%E0%B8%A5%E0%B8%B2%E0%B8%81%E0%B8%AA%E0%B8%B5%E0%B8%AA%E0%B8%B1%E0%B8%99%E0%B9%81%E0%B8%A5%E0%B8%B0%E0%B8%81%E0%B8%A3%E0%B8%B0%E0%B9%80%E0%B8%9B%E0%B9%8B%E0%B8%B2%E0%B9%80%E0%B8%9B%E0%B9%89%E0%B8%AA%E0%B8%B0%E0%B8%9E%E0%B8%B2%E0%B8%A2%E0%B8%AB%E0%B8%A5%E0%B8%B1%E0%B8%87%E0%B8%97%E0%B8%B5%E0%B9%88%E0%B8%88%E0%B8%B1%E0%B8%94%E0%B9%80%E0%B8%A3%E0%B8%B5%E0%B8%A2%E0%B8%87%E0%B8%9A%E0%B8%99%E0%B8%9E%E0%B8%B7%E0%B9%89%E0%B8%99%E0%B8%AB%E0%B8%A5%E0%B8%B1%E0%B8%87%E0%B8%AA%E0%B8%B5%E0%B8%99%E0%B9%89%E0%B9%8D%E0%B8%B2%E0%B9%80%E0%B8%87%E0%B8%B4%E0%B8%99.jpg?s=612x612&w=0&k=20&c=qSwdxfFLq4lV-e3j6bHXo2iH1gWqqHIM9Nyz9qnPlyc=" alt="Signup" /> 
            </div>
            <div className="signup-form">
                <h2>Create an Account</h2>
                <p>Enter your details below</p>
                <form onSubmit={handleSubmit}>
                    <input
                        type="text"
                        placeholder="Name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                        required
                    />
                    <input
                        type="email"
                        placeholder="Email "
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                    <input
                        type="password"
                        placeholder="Password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                    <button type="submit" className="signup-button">Create Account</button>
                </form>
                <p className="login-link">
                    Already have an account? <a href="/login">Log in</a>
                </p>
                <p>{message}</p>
            </div>
        </div>
    );
}

export default Register;
