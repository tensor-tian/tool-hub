// toolWorkerManager.ts - Manages WebWorker for tool plugin evaluation
import ToolWorkerUrl from './toolrc.worker?worker&url';

interface EvalToolRequest {
  type: 'eval-tool';
  code: string;
  parameters: string;
}

interface EvalToolResponse {
  type: 'eval-tool-result';
  success: boolean;
  tool?: any;
  error?: string;
  stack?: string;
}

class ToolWorkerManager {
  private worker: Worker | null = null;
  private isReady = false;
  private readyPromise: Promise<void>;
  private readyResolve!: () => void;

  constructor() {
    this.readyPromise = new Promise((resolve) => {
      this.readyResolve = resolve;
    });
    this.initWorker();
  }

  private initWorker() {
    this.worker = new Worker(ToolWorkerUrl, { type: 'module' });

    this.worker.onmessage = (e) => {
      if (e.data.type === 'ready') {
        this.isReady = true;
        this.readyResolve();
      }
    };

    this.worker.onerror = (err) => {
      console.error('ToolWorker error:', err);
    };
  }

  async evalTool(
    code: string,
    parameters: string,
  ): Promise<{ success: boolean; tool?: any; error?: string }> {
    if (!this.isReady) {
      await this.readyPromise;
    }

    return new Promise((resolve) => {
      if (!this.worker) {
        resolve({ success: false, error: 'Worker not initialized' });
        return;
      }

      const handler = (e: MessageEvent) => {
        const response: EvalToolResponse = e.data;
        if (response.type === 'eval-tool-result') {
          this.worker?.removeEventListener('message', handler);
          resolve({
            success: response.success,
            tool: response.tool,
            error: response.error,
          });
        }
      };

      this.worker.addEventListener('message', handler);

      const request: EvalToolRequest = {
        type: 'eval-tool',
        code,
        parameters,
      };

      this.worker.postMessage(request);
    });
  }

  destroy() {
    if (this.worker) {
      this.worker.terminate();
      this.worker = null;
      this.isReady = false;
    }
  }
}

// Export singleton instance
export const toolWorkerManager = new ToolWorkerManager();
