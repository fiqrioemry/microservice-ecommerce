// src/api/cart.js

import { cartInstance } from ".";

// =====================
// CART (Auth Required)
// =====================

const getCart = async () => {
  const res = await cartInstance.get("/cart");
  return res.data.cart;
};

const addToCart = async (data) => {
  const res = await cartInstance.post("/cart", data);
  return res.data;
};

const updateCartItem = async (itemId, data) => {
  const res = await cartInstance.put(`/cart/items/${itemId}`, data);
  return res.data;
};

const removeCartItem = async (itemId) => {
  const res = await cartInstance.delete(`/cart/items/${itemId}`);
  return res.data;
};

const clearCart = async () => {
  const res = await cartInstance.delete("/cart");
  return res.data;
};

export default {
  getCart,
  addToCart,
  updateCartItem,
  removeCartItem,
  clearCart,
};
