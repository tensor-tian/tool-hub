// toolEvalService.ts - Handles Wails events for tool evaluation
import { EventsOn, EventsEmit } from '../wailsjs/runtime/runtime';
import { toolWorkerManager } from './toolWorkerManager';

interface EvalToolRequestEvent {
  requestId: string;
  code: string;
  parameters: string;
}

interface EvalToolResponseEvent {
  requestId: string;
  success: boolean;
  tool?: any;
  error?: string;
}

export function initToolEvalService() {
  // Listen for eval-tool-request events from backend
  EventsOn('eval-tool-request', async (data: EvalToolRequestEvent) => {
    console.log('Received eval-tool-request:', data.requestId);

    try {
      const result = await toolWorkerManager.evalTool(data.code, data.parameters);

      const response: EvalToolResponseEvent = {
        requestId: data.requestId,
        success: result.success,
        tool: result.tool,
        error: result.error,
      };

      // Emit response back to backend
      await EventsEmit('eval-tool-response', response);
      console.log('Sent eval-tool-response:', data.requestId);
    } catch (err) {
      const response: EvalToolResponseEvent = {
        requestId: data.requestId,
        success: false,
        error: err instanceof Error ? err.message : String(err),
      };

      await EventsEmit('eval-tool-response', response);
      console.error('Error evaluating tool:', err);
    }
  });

  console.log('ToolEvalService initialized');
}
