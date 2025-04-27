// src/components/input/UploadElement.jsx
import { toast } from "sonner";
import { useState } from "react";
import { PlusCircle, X } from "lucide-react";
import { Controller, useFormContext } from "react-hook-form";

const UploadElement = ({
  name,
  label,
  maxImages = 5,
  maxSizeMB = 2,
  rules = { required: true },
}) => {
  const { control } = useFormContext();
  const [isDragging, setIsDragging] = useState(false);

  return (
    <Controller
      name={name}
      control={control}
      rules={rules}
      render={({ field, fieldState }) => {
        const handleFiles = (files) => {
          const fileArray = Array.from(files);
          const currentFiles = field.value || [];
          const validFiles = fileArray.filter((file) => {
            const isValidSize = file.size / (1024 * 1024) <= maxSizeMB;
            if (!isValidSize) {
              toast.warning(
                `${file.name} exceeds maximum size of ${maxSizeMB}MB`
              );
            }
            return isValidSize;
          });

          const updated = [...currentFiles, ...validFiles];

          if (updated.length > maxImages) {
            toast.warning(`You can only upload up to ${maxImages} images`);
            return;
          }

          field.onChange(updated);
        };

        const handleDrop = (e) => {
          e.preventDefault();
          setIsDragging(false);
          if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
            handleFiles(e.dataTransfer.files);
            e.dataTransfer.clearData();
          }
        };

        const handleDragOver = (e) => {
          e.preventDefault();
          setIsDragging(true);
        };

        const handleDragLeave = (e) => {
          e.preventDefault();
          setIsDragging(false);
        };

        const handleRemoveImage = (img) => {
          const updated = (field.value || []).filter((file) => file !== img);
          field.onChange(updated);
        };

        return (
          <div className="space-y-2">
            {label && (
              <label className="block text-sm font-medium text-gray-700">
                {label}
              </label>
            )}

            <div
              onDrop={handleDrop}
              onDragOver={handleDragOver}
              onDragLeave={handleDragLeave}
              className={`flex flex-wrap gap-4 p-4 border-2 ${
                isDragging
                  ? "border-primary bg-primary/10"
                  : "border-dashed border-primary"
              } rounded-md transition`}
            >
              {(field.value || []).map((img, idx) => (
                <div
                  key={idx}
                  className="relative w-32 h-32 rounded-md overflow-hidden border border-border"
                >
                  <img
                    src={URL.createObjectURL(img)}
                    alt="preview"
                    className="object-cover w-full h-full"
                  />
                  <button
                    type="button"
                    className="absolute top-1 right-1 bg-white rounded-full p-1 shadow-md hover:bg-red-500 hover:text-white transition"
                    onClick={() => handleRemoveImage(img)}
                  >
                    <X className="w-4 h-4" />
                  </button>
                </div>
              ))}

              {(!field.value || field.value.length < maxImages) && (
                <>
                  <label
                    htmlFor={`${name}-upload`}
                    className="flex flex-col items-center justify-center w-32 h-32 border-2 border-dashed border-primary rounded-md cursor-pointer hover:bg-primary/10 transition"
                  >
                    <PlusCircle className="text-primary mb-2" />
                    <span className="text-primary text-sm">Upload</span>
                  </label>
                  <input
                    id={`${name}-upload`}
                    type="file"
                    accept="image/*"
                    multiple
                    onChange={(e) => handleFiles(e.target.files)}
                    className="hidden"
                  />
                </>
              )}
            </div>

            {fieldState.error && (
              <p className="text-red-500 text-xs mt-1">
                {fieldState.error.message || "This field is required"}
              </p>
            )}
          </div>
        );
      }}
    />
  );
};

export { UploadElement };
