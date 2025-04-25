// src/pages/CartPage.jsx
import { useEffect } from "react";
import { useCartStore } from "@/store/useCartStore";
import { formatRupiah } from "@/lib/utils";
import Button from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Minus, Plus, Trash2 } from "lucide-react";

const CartPage = () => {
  const navigate = useNavigate();
  const {
    carts,
    fetchCart,
    totalPrice,
    totalItems,
    loading,
    updateItem,
    removeItem,
  } = useCartStore();

  useEffect(() => {
    fetchCart();
  }, []);

  const handleQuantityChange = (item, delta) => {
    const newQty = item.quantity + delta;
    if (newQty >= 1) {
      updateItem(item.id, { quantity: newQty });
    }
  };

  const isEmpty = carts.items?.length === 0;

  return (
    <div className="container mx-auto px-4 py-8 min-h-screen">
      <h1 className="text-2xl font-semibold mb-6">Shopping Cart</h1>

      {loading ? (
        <div className="text-center text-muted-foreground">Loading cart...</div>
      ) : isEmpty ? (
        <div className="flex flex-col items-center justify-center mt-16">
          <img src="/empty-cart.webp" alt="Empty Cart" className="w-40 h-40" />
          <p className="text-muted-foreground mt-4">Your cart is empty.</p>
        </div>
      ) : (
        <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
          {/* Cart Items */}
          <div className="lg:col-span-2 space-y-6">
            {carts.items.map((item) => (
              <div
                key={item.id}
                className="flex items-center gap-4 border p-4 rounded-lg shadow-sm"
              >
                <input
                  type="checkbox"
                  checked={item.isChecked}
                  onChange={() =>
                    updateItem(item.id, { isChecked: !item.isChecked })
                  }
                  className="w-5 h-5 accent-primary"
                />
                <img
                  src={item.imageUrl}
                  alt={item.name}
                  className="w-20 h-20 rounded object-cover border"
                />
                <div className="flex-1 text-sm">
                  <h2 className="font-semibold text-base">{item.name}</h2>
                  <p className="text-muted-foreground mt-1">
                    Harga: {formatRupiah(item.price)}
                  </p>
                  <p className="text-muted-foreground mt-1">
                    Subtotal: {formatRupiah(item.price * item.quantity)}
                  </p>
                  <div className="flex items-center gap-2 mt-2">
                    <button
                      onClick={() => handleQuantityChange(item, -1)}
                      className="px-2 py-1 border rounded"
                    >
                      <Minus className="w-4 h-4" />
                    </button>
                    <span className="w-6 text-center">{item.quantity}</span>
                    <button
                      onClick={() => handleQuantityChange(item, 1)}
                      className="px-2 py-1 border rounded"
                    >
                      <Plus className="w-4 h-4" />
                    </button>
                  </div>
                </div>
                <button
                  onClick={() => removeItem(item.id)}
                  className="text-red-500 hover:text-red-700 ml-2"
                  title="Remove item"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>
            ))}
          </div>

          {/* Summary */}
          <div className="border p-6 rounded-lg shadow-sm">
            <h2 className="text-lg font-semibold mb-4">Order Summary</h2>
            <div className="flex justify-between mb-2">
              <span>Total Items</span>
              <span>{totalItems}</span>
            </div>
            <div className="flex justify-between font-medium text-lg mb-4">
              <span>Total</span>
              <span>{formatRupiah(totalPrice)}</span>
            </div>
            <Button className="w-full" onClick={() => navigate("/checkout")}>
              Proceed to Checkout
            </Button>
          </div>
        </div>
      )}
    </div>
  );
};

export default CartPage;
