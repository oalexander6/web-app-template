import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

export default function Hello() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ['hellomsg'],
    queryFn: () => axios.get('/api/v1').then((res) => res.data)
  });

  return (
    <>
      <p className="text-white">Pending: {isPending ? 'true' : 'false'}</p>
      <p className="text-white">Is Error: {isError ? 'true' : 'false'}</p>
      <p className="text-white">Error: {error?.message}</p>
      <p className="text-white">Data: {data}</p>
    </>
  );
}
