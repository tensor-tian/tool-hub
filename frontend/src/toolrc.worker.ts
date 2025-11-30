// toolrc.worker.ts - WebWorker script for toolrc plugin
import { z } from 'zod';
import { createAuxiliaryTypeStore, zodToTs, printNode, createTypeAlias } from 'zod-to-ts';
import { zerialize, type ZodTypes } from 'zodex';

// Helper functions to pass to plugin
const toJSONSchema = (s: z.ZodTypeAny): string => {
  return JSON.stringify(z.toJSONSchema(s), null, 2);
};

const toTSDefinition = (name: string, s: z.ZodTypeAny): string => {
  const auxiliaryTypeStore = createAuxiliaryTypeStore();
  const { node } = zodToTs(s, { auxiliaryTypeStore });
  const typeAlias = createTypeAlias(
    node,
    `${name.charAt(0).toUpperCase() + name.slice(1)}Parameters`,
  );
  return printNode(typeAlias);
};

const serializeZod = (s: z.ZodTypeAny): string => {
  return JSON.stringify(zerialize(s as ZodTypes), null, 2);
};

const deps = {
  z,
  toJSONSchema,
  toTSDefinition,
  serializeZod,
};

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

// Load dependencies and set up message handler
self.onmessage = (e) => {
  const request: EvalToolRequest = e.data;

  if (request.type !== 'eval-tool') {
    self.postMessage({
      type: 'eval-tool-result',
      success: false,
      error: 'Invalid request type',
    } as EvalToolResponse);
    return;
  }

  const { code, parameters } = request;

  try {
    // Execute the plugin code
    const fn = new Function(`${code}; return ToolPlugin;`);
    const plugin = fn();

    // Create tool with all dependencies
    const toolFactory = plugin.defineTool(deps);

    // Parse parameters
    const params = JSON.parse(parameters);

    // Create tool with provided parameters
    const tool = toolFactory.createTool(params);

    self.postMessage({
      type: 'eval-tool-result',
      success: true,
      tool,
    } as EvalToolResponse);
  } catch (err) {
    self.postMessage({
      type: 'eval-tool-result',
      success: false,
      error: (err as Error).message,
      stack: (err as Error).stack,
    } as EvalToolResponse);
  }
};

// Signal that worker is ready
self.postMessage({ type: 'ready' });
