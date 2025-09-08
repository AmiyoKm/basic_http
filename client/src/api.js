import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL || "/api/products";

const api = axios.create({
	baseURL: API_URL.includes("http") ? "" : window.location.origin,
	headers: {
		amiyo: "secret",
	},
});

export const getProducts = () => api.get(API_URL);

export const createProduct = (product) => api.post(API_URL, product);

export const updateProduct = (id, product) =>
	api.put(`${API_URL}/${id}`, product);

export const deleteProduct = (id) => api.delete(`${API_URL}/${id}`);
