
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './css/CategoryProducts.css';

const CategoryProducts = () => {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const fetchProducts = async () => {
            try {
                const response = await axios.get('/api/v1/products/category/2');
                setProducts(response.data);
                setLoading(false);
            } catch (err) {
                setError(err.message);
                setLoading(false);
            }
        };

        fetchProducts();
    }, []);

    if (loading) return <p className="loading">Loading...</p>;
    if (error) return <p className="error">Error: {error}</p>;

    return (
        <div className="product-list">
            <h1> ดินสอ/ปากกา</h1>
            <div className="product-cards">
                {products.map((product) => (
                    <div key={product.id} className="product-card">
                        <img 
                            src={product.images[0]?.image_url || "path-to-placeholder-image.jpg"} 
                            alt={product.name} 
                            className="product-image" 
                        />
                        <h5>{product.name}</h5>
                        <p className="product-price">฿{product.price}</p>
                        <p className="product-seller">
                            {product.inventory.quantity > 0 ? 'In stock' : 'Out of stock'}
                        </p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default CategoryProducts;
