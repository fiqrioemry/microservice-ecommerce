import React from "react";
import ErrorDialog from "@/components/ui/ErrorDialog";
import FetchLoading from "@/components/ui/FetchLoading";
import { Card, CardContent } from "@/components/ui/card";
import { ScrollArea } from "@/components/ui/scroll-area";
import AddAddress from "@/components/address/AddAddress";
import { useAddressesQuery } from "@/hooks/useProfileManagement";
import UpdateAddress from "../../components/address/UpdateAddress";
// -- 2,24,4,711,4
const Address = () => {
  const { isError, isLoading, data: addresses } = useAddressesQuery();

  if (isLoading) return <FetchLoading />;

  if (isError) return <ErrorDialog />;

  return (
    <div className="space-y-4">
      <div className="flex justify-between items-center">
        <h2 className="text-lg font-semibold">My Addresses</h2>
        <AddAddress />
      </div>

      <ScrollArea className="h-[500px] space-y-4">
        {addresses.addresses.map((add) => (
          <Card key={add.id}>
            <CardContent className="p-4 space-y-2">
              <div className="flex items-center justify-between">
                <h4 className="text-sm font-medium capitalize">{add.name}</h4>
                {add.isMain && (
                  <span className="text-xs text-white bg-primary rounded px-2">
                    Main
                  </span>
                )}
              </div>
              <p className="text-sm text-muted-foreground">{add.phone}</p>
              <p className="text-sm">{add.address}</p>
              <p className="text-sm">
                {add.province}, {add.city}, {add.district}, {add.subdistrict},{" "}
                {add.postalCode}
              </p>
              <div className="flex gap-2 text-sm text-primary underline">
                <UpdateAddress address={add} />
                <button>Delete</button>
              </div>
            </CardContent>
          </Card>
        ))}
      </ScrollArea>
    </div>
  );
};

export default Address;
