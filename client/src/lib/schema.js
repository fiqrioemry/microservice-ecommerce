import { z } from "zod";

export const signInSchema = z.object({
  email: z.string().email("Email tidak valid"),
  password: z.string().min(6, "Password minimal 6 karakter"),
  rememberMe: z.boolean().optional(),
});

export const registerSchema = z.object({
  email: z.string().email("Email tidak valid"),
  fullname: z.string().min(6, "Nama lengkap minimal 6 karakter"),
  password: z.string().min(6, "Password minimal 6 karakter"),
});

export const addressSchema = z.object({
  name: z.string().min(6, "Nama minimal 6 karakter"),
  address: z.string().min(1, "Alamat tidak boleh kosong"),
  provinceId: z.string().min(1, "Province harus dipilih"),
  cityId: z.string().min(1, "City harus dipilih"),
  districtId: z.string().min(1, "District harus dipilih"),
  subdistrictId: z.string().min(1, "Subdistrict harus dipilih"),
  postalCodeId: z.string().min(1, "Postal code harus dipilih"),
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
  imageUrl: z.string().url().optional(), // URL dari hasil upload gambar variant
  options: z.record(z.string()).optional(), // Map {typeName: value}
});

// Attribute Schema
const attributeSchema = z.object({
  attributeId: z.number().int(),
  attributeValueId: z.number().int(),
});

export const createProductSchema = z.object({
  name: z.string().min(1, "Product name is required"),
  description: z.string().optional(),
  categoryId: z.string().uuid({ message: "Invalid category ID" }),
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
        .refine((file) => file.size <= 5 * 1024 * 1024, {
          message: "Max 5MB image size",
        })
    )
    .min(1, { message: "At least 1 product image is required" }),
  variants: z
    .array(variantSchema)
    .min(1, { message: "At least 1 variant is required" }),
  attributes: z.array(attributeSchema).optional(),
});
