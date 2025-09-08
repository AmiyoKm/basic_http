
import React, { useState, useEffect } from 'react';

const ProductForm = ({ product, onSave, onCancel }) => {
  const [formData, setFormData] = useState({
    name: '',
    description: '',
    imageUrl: '',
    price: 0,
  });

  useEffect(() => {
    if (product) {
      setFormData(product);
    } else {
      setFormData({
        name: '',
        description: '',
        imageUrl: '',
        price: 0,
      });
    }
  }, [product]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevData => ({
      ...prevData,
      [name]: name === 'price' ? (value === '' ? '' : parseFloat(value)) : value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave(formData);
  };

  return (
    <div className="fixed inset-0 bg-gray-800 bg-opacity-50 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-md w-full max-w-md">
        <h2 className="text-2xl font-bold mb-4">{product ? 'Edit Product' : 'Add Product'}</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-700">Name</label>
            <input type="text" name="name" value={formData.name} onChange={handleChange} className="w-full px-3 py-2 border rounded-md" required />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Description</label>
            <textarea name="description" value={formData.description} onChange={handleChange} className="w-full px-3 py-2 border rounded-md" required></textarea>
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Image URL</label>
            <input type="text" name="imageUrl" value={formData.imageUrl} onChange={handleChange} className="w-full px-3 py-2 border rounded-md" required />
          </div>
          <div className="mb-4">
            <label className="block text-gray-700">Price</label>
            <input type="number" name="price" value={formData.price} onChange={handleChange} className="w-full px-3 py-2 border rounded-md" required />
          </div>
          <div className="flex justify-end">
            <button type="button" onClick={onCancel} className="bg-gray-500 text-white px-4 py-2 rounded-md mr-2">Cancel</button>
            <button type="submit" className="bg-blue-500 text-white px-4 py-2 rounded-md">Save</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default ProductForm;
