import React from 'react';
import { createRoot } from 'react-dom/client';
import '@/global.css';
import App from '@/App';
import { initToolEvalService } from '@/toolEvalService';

const container = document.getElementById('root');

const root = createRoot(container!);

// Initialize tool evaluation service
initToolEvalService();

root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
