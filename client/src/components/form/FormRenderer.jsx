import {
  Select,
  SelectTrigger,
  SelectContent,
  SelectItem,
  SelectValue,
} from "@/components/ui/select";
import { Label } from "@/components/ui/label";
import { Controller } from "react-hook-form";
import { Input } from "@/components/ui/input";

export const FormRenderer = ({ control, schema }) => {
  return schema.map((field) => (
    <Controller
      key={field.name}
      name={field.name}
      control={control}
      render={({ field: rhfField, fieldState }) => (
        <div className="space-y-2 col-span-2">
          <Label>{field.label}</Label>
          {field.type === "input" && (
            <Input
              {...rhfField}
              type={field.inputType || "text"}
              placeholder={field.placeholder}
            />
          )}
          {field.type === "select" && (
            <Select value={rhfField.value} onValueChange={rhfField.onChange}>
              <SelectTrigger>
                <SelectValue placeholder={field.placeholder} />
              </SelectTrigger>
              <SelectContent>
                {field.options?.map((opt) => (
                  <SelectItem key={opt.value} value={opt.value}>
                    {opt.label}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          )}
          {fieldState.error && (
            <p className="text-sm text-red-500">{fieldState.error.message}</p>
          )}
        </div>
      )}
    />
  ));
};
