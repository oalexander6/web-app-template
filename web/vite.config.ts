import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import { config as dotenvConfig } from 'dotenv';

dotenvConfig({
    path: '../.env'
});

const serverPort = process.env.PORT;

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
    server: {
        proxy: {
            '/api': {
                target: `http://localhost:${serverPort}/`,
                changeOrigin: true
            }
        }
    }
});
