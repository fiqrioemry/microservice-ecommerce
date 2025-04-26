import { authInstance, publicInstance } from ".";
const errorHandle = (error) => {
  const errorMessage = error.response?.data?.message;
  return Promise.reject(new Error(errorMessage));
};

const callApi = {
  // NOTE : PUBLIC API route
  // User Authentication management
  login: async (formData) => {
    return publicInstance
      .post("/auth/login", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  logout: async () => {
    return authInstance
      .post("/auth/logout")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  register: async (formData) => {
    return publicInstance
      .post("/auth/register", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  sendOTP: async (formData) => {
    return publicInstance
      .post("/auth/send-otp", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  verifyOTP: async (formData) => {
    return publicInstance
      .post("/auth/verify-otp", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  refreshToken: async () => {
    return publicInstance
      .post("/auth/refresh")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  authCheck: async () => {
    return authInstance
      .get("/auth/me")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // products management
  getProduct: async (slug) => {
    return authInstance
      .get(`/product/${slug}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getProducts: async (searchParams) => {
    const queryString = new URLSearchParams(searchParams).toString();
    return authInstance
      .get(`/product?${queryString}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  searchProducts: async (search) => {
    return authInstance
      .get(`/product?search=${search}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getCategories: async () => {
    return authInstance
      .get("/category")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getStoreInfo: async (shopname) => {
    return authInstance
      .get(`/store/${shopname}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // note : CUSTOMER API Route
  // customer profile management
  getProfile: async () => {
    return authInstance
      .get("/user/profile")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateProfile: async (formData) => {
    return authInstance
      .put("/user/profile", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getAddress: async () => {
    return authInstance
      .get("/user/profile/address")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  addAddress: async (formData) => {
    return authInstance
      .post("/user/profile/address", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateAddress: async (formData, addressId) => {
    return authInstance
      .put(`/user/profile/address/${addressId}`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  deleteAddress: async (addressId) => {
    return authInstance
      .delete(`/user/profile/address/${addressId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // customer transactions and orders management
  getAllUserOrders: async (params) => {
    return authInstance
      .get(`/customer/orders?status=${params}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getUserOrderDetail: async (orderId) => {
    return authInstance
      .get(`/customer/orders/${orderId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getShipmentDetail: async (orderId) => {
    return authInstance
      .get(`/customer/orders/${orderId}/shipment`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  cancelTransaction: async (formData, transactionId) => {
    return authInstance
      .put(`/customer/transactions/${transactionId}/cancel`, { formData })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  cancelUserOrder: async (formData, orderId) => {
    return authInstance
      .put(`/customer/orders/${orderId}/cancel`, { formData })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  confirmOrderDelivery: async (formData, orderId) => {
    console.log(formData);
    return authInstance
      .put(`/customer/orders/${orderId}/confirm`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getAllTransactions: async (params) => {
    console.log(params);
    return authInstance
      .get(`/customer/transactions?status=${params}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getTransactionDetail: async (transactionId) => {
    return authInstance
      .get(`/customer/transactions/${transactionId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  createNewTransaction: async (transactionData) => {
    return authInstance
      .post("/customer/transactions", transactionData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getUserNotifications: async () => {
    return authInstance
      .get(`/customer/notifications`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // customer carts management
  getCarts: async () => {
    return authInstance
      .get("/cart")
      .then((res) => res.data)
      .catch(errorHandle);
  },

  addCart: async (productId, quantity) => {
    return authInstance
      .post("/cart", { productId, quantity })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateCart: async (cartId, quantity) => {
    return authInstance
      .put(`/cart/${cartId}`, { quantity })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  removeCart: async (cartId) => {
    return authInstance
      .delete(`/cart/${cartId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // note : ADMIN API Route
  // category management
  createNewCategory: async (formData) => {
    return authInstance
      .post(`/category`, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateCategory: async (formData, categoryId) => {
    return authInstance
      .put(`/category/${categoryId}`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  deleteCategory: async (categoryId) => {
    return authInstance
      .delete(`/category/${categoryId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getAdminDashboardSummary: async () => {
    return authInstance
      .get(`/admin/statistic`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getAllUsers: async () => {
    return authInstance
      .get(`/admin/users`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getAllShipments: async () => {
    return authInstance
      .get(`/admin/shipments`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateShipmentStatus: async (status, shipmentId) => {
    return authInstance
      .put(`/admin/shipments/${shipmentId}`, status)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // note : SELLER  API Route
  // open store management
  createStore: async (formData) => {
    return authInstance
      .post("/auth/open-store", formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // store profile management
  getStoreProfile: async () => {
    return authInstance
      .get(`/store`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // TODO : Create feature update store profile
  updateStoreProfile: async (formData) => {
    return authInstance
      .put(`/store`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // store statistic management
  getStoreStatisticSummary: async (formData) => {
    return authInstance
      .get(`/seller/statistic`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getStoreNotifications: async () => {
    return authInstance
      .get(`/seller/notifications`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // store order management
  getAllStoreOrders: async (params) => {
    return authInstance
      .get(`/seller/orders?status=${params}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  getStoreOrderDetail: async (orderId) => {
    return authInstance
      .get(`/seller/orders/${orderId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  cancelStoreOrder: async (formData, orderId) => {
    return authInstance
      .put(`/seller/orders/${orderId}/cancel`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  proceedStoreOrder: async (formData, orderId) => {
    return authInstance
      .put(`/seller/orders/${orderId}/process`, formData)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  // store products management
  getStoreProducts: async (searchParams) => {
    const queryString = new URLSearchParams(searchParams).toString();
    return authInstance
      .get(`/store/product?${queryString}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },

  createProduct: async (formData) => {
    return authInstance
      .post("/store/product", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  updateProduct: async (formData, productId) => {
    return authInstance
      .put(`/store/product/${productId}`, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      })
      .then((res) => res.data)
      .catch(errorHandle);
  },

  deleteProduct: async (productId) => {
    return authInstance
      .delete(`/store/product/${productId}`)
      .then((res) => res.data)
      .catch(errorHandle);
  },
};

export default callApi;
