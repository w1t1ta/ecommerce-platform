// Contact.jsx
import React from 'react';
import { Link } from 'react-router-dom';

import './css/Contact.css';

const Contact = () => {
    return (
        <div className="contact-page">
            <div className="contact-container">
                {/* Contact Information Section */}
                <div className="contact-info">
                    <div className="contact-item">
                        <h2>Call To Us</h2>
                        <p>We are available 24/7, 7 days a week.</p>
                        <p>Phone: 123-123-12121</p>
                    </div>
                    <div className="contact-item">
                        <h2>Write To Us</h2>
                        <p>Fill out our form and we will contact you within 24 hours.</p>
                        <p>Email: customer@exclusive.com</p>
                        <p>Email: support@exclusive.com</p>
                    </div>
                </div>
                {/* Contact Form Section */}
                <div className="contact-form">
                    <form>
                        <div className="form-group">
                            <input type="text" placeholder="Your Name *" required />
                            <input type="email" placeholder="Your Email *" required />
                            <input type="tel" placeholder="Your Phone *" required />
                        </div>
                        <div className="form-group">
                            <textarea placeholder="Your Message" rows="6" required></textarea>
                        </div>
                        <button type="submit" className="submit-button">Send Message</button>
                    </form>
                </div>
            </div>
        </div>
    );
};

export default Contact;
