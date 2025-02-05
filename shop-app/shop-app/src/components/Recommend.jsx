// src/components/Recommend.jsx
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import './css/Recommend.css';

const Recommend = () => {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
           
        axios.get('api/v1/products/recommend')
            .then((response) => {
                setProducts(response.data);
                setLoading(false);
            })
            .catch((err) => {
                setError(err.message);
                setLoading(false);
            });
    }, []);

    if (loading) return <div className="loading">Loading...</div>;
    if (error) return <div className="error">Error: {error}</div>;

    return (
        <div className="recommend-container">
            <h1>สินค้าแนะนำ</h1>
            <div className="recommend-list">
                {products.map((product) => (
                    <div key={product.id} className="recommend-card">
                        {/* Display product image if available, otherwise use placeholder */}
                        
                        <img 
                            src={product.images?.[0]?.image_url || "path-to-placeholder-image.jpg"} 

                            alt={product.name} 
                            className="product-image" 
                        />
                        
                        <h3 className="product-name">{product.name}</h3>
                        {/* <p className="product-description">{product.description}</p> */}
                        <p className="product-price">ราคา: {product.price} บาท</p>

                        <Link to={`/product/${product.id}`} className="view-details">
                            View Details
                        </Link>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default Recommend;




