import { getProducts } from "@/lib/api";

export default async function Page() {
  const products = await getProducts();

  return (
    <div>
      <h1>Products</h1>

      {products.map((p: any) => (
        <div key={p.id}>
          {p.name}
        </div>
      ))}
    </div>
  );
}