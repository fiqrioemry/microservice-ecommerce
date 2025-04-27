// src/lib/constant.js
export const addressState = {
  name: "",
  address: "",
  provinceId: "",
  districtId: "",
  subdistrictId: "",
  postalCodeId: "",
  cityId: "",
  phone: "",
  isMain: false,
};

export const registerState = {
  fullname: "",
  password: "",
  email: "",
};

export const signInState = {
  password: "",
  email: "",
  rememberMe: false,
};

export const createProductState = {
  name: "",
  description: "",
  categoryId: "",
  subcategoryId: "",
  isFeatured: false,
  weight: 0,
  length: 0,
  width: 0,
  height: 0,
  discount: 0,
  images: [],
  variants: [
    {
      sku: "",
      price: 0,
      stock: 0,
      sold: 0,
      isActive: true,
      imageUrl: "",
      options: {},
    },
  ],
  attributes: [],
};
