import { useQuery } from '@tanstack/react-query';
import axios from 'axios';

import { ThemeToggle } from './theme-toggle';

export default function Hello() {
  const { isPending, isError, data, error } = useQuery({
    queryKey: ['hellomsg'],
    queryFn: () => axios.get('/api/v1').then((res) => res.data)
  });

  return (
    <>
      <ThemeToggle />
      <p className="dark:text-white">Pending: {isPending ? 'true' : 'false'}</p>
      <p className="dark:text-white">Is Error: {isError ? 'true' : 'false'}</p>
      <p className="dark:text-white">Error: {error?.message}</p>
      <p className="dark:text-white">Data: {data}</p>
    </>
  );
}
