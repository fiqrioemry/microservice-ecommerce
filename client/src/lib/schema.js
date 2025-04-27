// src/lib/schema.js
import { z } from "zod";

export const signInSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
  rememberMe: z.boolean().optional(),
});

export const registerSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
  fullname: z.string().min(6, "Nama lengkap minimal 6 karakter"),
});

export const addressSchema = z.object({
  cityId: z.string().min(1, "City harus dipilih"),
  name: z.string().min(6, "Nama minimal 6 karakter"),
  address: z.string().min(1, "Alamat tidak boleh kosong"),
  districtId: z.string().min(1, "District harus dipilih"),
  provinceId: z.string().min(1, "Province harus dipilih"),
  postalCodeId: z.string().min(1, "Postal code harus dipilih"),
  subdistrictId: z.string().min(1, "Subdistrict harus dipilih"),
  phone: z
    .string()
    .min(10, "Nomor telepon minimal 10 karakter")
    .max(12, "Nomor telepon maksimal 12 karakter"),
  isMain: z.boolean().optional(),
});

// Variant Schema
const variantSchema = z.object({
  sku: z.string().min(1, "SKU is required"),
  price: z
    .number({ required_error: "Price is required" })
    .positive("Price must be greater than 0"),
  stock: z
    .number({ required_error: "Stock is required" })
    .int()
    .nonnegative("Stock cannot be negative"),
  sold: z.number().int().nonnegative().default(0),
  isActive: z.boolean().default(true),
  imageUrl: z.string().url().optional(),
  options: z.record(z.string()).optional(),
});

const attributeSchema = z.object({
  attributeId: z.number().int(),
  attributeValueId: z.number().int(),
});

export const createProductSchema = z.object({
  name: z.string().min(20, "Minimal nama product 20 arakter"),
  description: z.string().optional(),
  categoryId: z
    .string()
    .uuid({ message: "Invalid category ID" })
    .min(1, "Category wajib diisi"),
  subcategoryId: z
    .string()
    .uuid({ message: "Invalid subcategory ID" })
    .optional(),
  isFeatured: z.boolean().default(false),
  weight: z.number().nonnegative().default(0),
  length: z.number().nonnegative().default(0),
  width: z.number().nonnegative().default(0),
  height: z.number().nonnegative().default(0),
  discount: z.number().nonnegative().optional(),
  images: z
    .array(
      z
        .instanceof(File)
        .refine((file) => file.type.startsWith("image/"), {
          message: "File must be an image",
        })
        .refine((file) => file.size <= 2 * 1024 * 1024, {
          message: "Max 2MB image size",
        })
    )
    .min(1, { message: "At least 1 product image is required" }),
  variants: z
    .array(variantSchema)
    .min(1, { message: "At least 1 variant is required" }),
  attributes: z.array(attributeSchema).optional(),
});

// Category Schema
export const categorySchema = z.object({
  name: z.string().min(1, "Name is required"),
  image: z
    .array(
      z
        .instanceof(File)
        .refine((file) => file.type.startsWith("image/"), {
          message: "File must be an image",
        })
        .refine((file) => file.size <= 2 * 1024 * 1024, {
          message: "Max 2MB image size",
        })
    )
    .min(1, { message: "At least 1 product image is required" }),
});

//  Subcategory Schema
export const subCategorySchema = z.object({
  name: z.string().min(1, "Name is required"),
  image: z
    .array(
      z
        .instanceof(File)
        .refine((file) => file.type.startsWith("image/"), {
          message: "File must be an image",
        })
        .refine((file) => file.size <= 2 * 1024 * 1024, {
          message: "Max 2MB image size",
        })
    )
    .min(1, { message: "At least 1 product image is required" }),
  categoryId: z.string().uuid({ message: "Invalid category ID" }),
});
