// src/category.ts
import { productInstance } from ".";

const getBanners = async (position) => {
  const res = await productInstance.get(`/banners/${position}`);
  return res.data;
};

export default {
  getBanners,
};
