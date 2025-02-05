import React, { useState } from 'react';
import axios from 'axios';
import Header from '../components/Header'; // Ensure the path is correct
import Footer from '../components/Footer'; // Ensure the path is correct
const API_URL = '/api/v1/shipping-addresses';

function CheckOut() {
    const [form, setForm] = useState({
        firstName: '',
        companyName: '',
        address: '',
        apartment: '',
        city: '',
        country: '',
        phone: '',
        email: '',
    });

    const [paymentMethod, setPaymentMethod] = useState("COD");

    const handleChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        });
    };

    const handlePaymentChange = (e) => {
        setPaymentMethod(e.target.value);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        try {
            // Make an API request to save the order (if required)
            // Example: await axios.post('/api/order', { form, paymentMethod });
            
            alert('Order Placed Successfully!');
        } catch (error) {
            console.error("Error placing order:", error);
            alert('Failed to place order. Please try again.');
        }
    };

    return (
        <div className="checkout-page">
            <Header />

            <div className="container">
                {/* Billing Details Form */}
                <div className="form-section">
                    <h2>Billing Details</h2>
                    <form onSubmit={handleSubmit}>
                        <label>First Name*</label>
                        <input
                            type="text"
                            name="firstName"
                            value={form.firstName}
                            onChange={handleChange}
                            placeholder="Enter your first name"
                            required
                        />

                        <label>Company Name</label>
                        <input
                            type="text"
                            name="companyName"
                            value={form.companyName}
                            onChange={handleChange}
                            placeholder="Enter company name"
                        />

                        <label>Street Address*</label>
                        <input
                            type="text"
                            name="address"
                            value={form.address}
                            onChange={handleChange}
                            placeholder="Enter street address"
                            required
                        />

                        <label>Apartment, suite, etc. (optional)</label>
                        <input
                            type="text"
                            name="apartment"
                            value={form.apartment}
                            onChange={handleChange}
                            placeholder="Enter apartment number"
                        />

                        <label>City*</label>
                        <input
                            type="text"
                            name="city"
                            value={form.city}
                            onChange={handleChange}
                            placeholder="Enter city"
                            required
                        />

                        <label>Country</label>
                        <input
                            type="text"
                            name="country"
                            value={form.country}
                            onChange={handleChange}
                            placeholder="Enter country"
                        />

                        <label>Phone Number*</label>
                        <input
                            type="text"
                            name="phone"
                            value={form.phone}
                            onChange={handleChange}
                            placeholder="Enter phone number"
                            required
                        />

                        <label>Email Address*</label>
                        <input
                            type="email"
                            name="email"
                            value={form.email}
                            onChange={handleChange}
                            placeholder="Enter email address"
                            required
                        />

                        <div>
                            <input type="checkbox" id="save-info" />
                            <label htmlFor="save-info">
                                Save this information for faster check-out next time
                            </label>
                        </div>
                    </form>
                </div>

                {/* Order Summary */}
                <div className="summary-section">
                    <h2>Order Summary</h2>
                    <img src="/image/devops.png"  width="45" height="45"></img>
                    <p>Practical DEVOPS and Cloud Engineering book - $545</p>
                    <img src="/image/canson.png"  width="45" height="45"></img>
                    <p>Canson สมุดวาดเขียนสีน้ำ A5 300 แกรม ชนิดเรียบ  - $106</p>
                    
                    <hr />
                    <p>Subtotal: $651</p>
                    <p>Shipping: Free</p>
                    <p>Total: $651</p>

                    {/* Payment Method */}
                    <div>
                        <input
                            type="radio"
                            id="cod"
                            name="payment"
                            value="COD"
                            checked={paymentMethod === "COD"}
                            onChange={handlePaymentChange}
                        />
                        <label htmlFor="cod">Cash on Delivery</label>

                        <input
                            type="radio"
                            id="bank"
                            name="payment"
                            value="Bank"
                            checked={paymentMethod === "Bank"}
                            onChange={handlePaymentChange}
                        />
                        <label htmlFor="bank">Bank</label>
                    </div>

                    <button
                        type="button"
                        className="place-order-btn"
                        onClick={handleSubmit}
                    >
                        Place Order
                    </button>
                </div>
            </div>

            <Footer />
        </div>
    );
}

export default CheckOut;
