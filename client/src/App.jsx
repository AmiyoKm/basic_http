import { useEffect, useState } from "react";
import {
    createProduct,
    deleteProduct,
    getProducts,
    updateProduct,
} from "./api";
import Header from "./components/Header";
import ProductForm from "./components/ProductForm";
import ProductList from "./components/ProductList";

const App = () => {
    const [products, setProducts] = useState([]);
    const [editingProduct, setEditingProduct] = useState(null);
    const [isFormVisible, setIsFormVisible] = useState(false);

    useEffect(() => {
        fetchProducts();
    }, []);

    const fetchProducts = async () => {
        try {
            const response = await getProducts();
            setProducts(response.data.value);
        } catch (error) {
            console.error("Error fetching products:", error);
        }
    };

    const handleSave = async (product) => {
        try {
            if (product.id) {
                await updateProduct(product.id, product);
            } else {
                await createProduct(product);
            }
            fetchProducts();
            setEditingProduct(null);
            setIsFormVisible(false);
        } catch (error) {
            console.error("Error saving product:", error);
        }
    };

    const handleEdit = (product) => {
        setEditingProduct(product);
        setIsFormVisible(true);
    };

    const handleDelete = async (id) => {
        try {
            await deleteProduct(id);
            fetchProducts();
        } catch (error) {
            console.error("Error deleting product:", error);
        }
    };

    const handleCancel = () => {
        setEditingProduct(null);
        setIsFormVisible(false);
    };

    const handleAdd = () => {
        setEditingProduct(null);
        setIsFormVisible(true);
    };

    return (
        <div className="bg-gray-100 min-h-screen">
            <Header />
            <main className="p-4">
                <div className="flex justify-end mb-4">
                    <button
                        onClick={handleAdd}
                        className="bg-green-500 text-white px-4 py-2 rounded-md"
                    >
                        Add Product
                    </button>
                </div>
                <ProductList
                    products={products}
                    onEdit={handleEdit}
                    onDelete={handleDelete}
                />
                {isFormVisible && (
                    <ProductForm
                        product={editingProduct}
                        onSave={handleSave}
                        onCancel={handleCancel}
                    />
                )}
            </main>
        </div>
    );
};

export default App;
