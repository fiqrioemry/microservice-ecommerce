// src/category.ts
import { productInstance } from ".";

const getAllBanners = async () => {
  const res = await productInstance.get(`/banners`);
  return res.data;
};

const getBannersByPosition = async (position) => {
  const res = await productInstance.get(`/banners/${position}`);
  return res.data;
};

export default {
  getAllBanners,
  getBannersByPosition,
};
