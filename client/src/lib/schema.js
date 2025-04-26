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
  address: z.string().min(1, "Address tidak boleh kosong"),
  provinceId: z.string().min(1, "Province harus dipilih"),
  cityId: z.string().min(1, "City harus dipilih"),
  zipcode: z.string().min(1, "Zipcode tidak boleh kosong"),
  phone: z.string().min(12, "Nomor telepon tidak valid"),
  isMain: z.boolean().optional(),
});
