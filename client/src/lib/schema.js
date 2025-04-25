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
