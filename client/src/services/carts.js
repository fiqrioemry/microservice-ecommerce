// src/api/cart.js
import { instance } from ".";

// =====================
// CART (Auth Required)
// =====================

const getCart = async () => {
  const res = await instance.get("/cart");
  return res.data;
};

const addToCart = async (data) => {
  const res = await instance.post("/cart", data);
  return res.data;
};

const updateCartItem = async (itemId, data) => {
  const res = await instance.put(`/cart/items/${itemId}`, data);
  return res.data;
};

const removeCartItem = async (itemId) => {
  const res = await instance.delete(`/cart/items/${itemId}`);
  return res.data;
};

const clearCart = async () => {
  const res = await instance.delete("/cart");
  return res.data;
};

export default {
  getCart,
  addToCart,
  updateCartItem,
  removeCartItem,
  clearCart,
};
