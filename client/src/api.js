import axios from "axios";

const API_URL_PRODUCTS = import.meta.env.VITE_API_URL || "/api/products";

const api = axios.create({
    baseURL: API_URL_PRODUCTS.includes("http") ? "" : window.location.origin,
	headers: {
		amiyo: "secret",
	},
});

export const getProducts = () => api.get(API_URL_PRODUCTS);

export const createProduct = (product) => api.post(API_URL_PRODUCTS, product);

export const updateProduct = (id, product) =>
	api.put(`${API_URL_PRODUCTS}/${id}`, product);

export const deleteProduct = (id) => api.delete(`${API_URL_PRODUCTS}/${id}`);
