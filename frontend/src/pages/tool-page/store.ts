import { create } from 'zustand';
import { immer } from 'zustand/middleware/immer';
import { combine } from 'zustand/middleware';
import { GetSettings, GetToolList, GetCommandLineTool } from '@/../wailsjs/go/hub/Model';
import { hub } from '@/../wailsjs/go/models';

export { useShallow } from 'zustand/shallow';

interface TToolTestcase {
  toolName: string;
  input: string;
  output: string;
  ok: boolean;
}

const SettingKeys = hub.StringValues;

export const useToolPageStore = create(
  immer(
    combine(
      {
        toolsDir: '',
        tools: [] as hub.Tool[],
        testcases: [] as TToolTestcase[],
        activeTool: null as hub.CommandLineTool | null,
      },
      (set, get) => {
        return {
          init: async () => {
            const res = await GetSettings([SettingKeys.SettingKeyToolsDir]);
            if (res.error.length === 0) {
              set((d) => {
                d.toolsDir = res.kvMap[SettingKeys.SettingKeyToolsDir] || '';
              });
            }
            set((d) => {
              d.tools = [];
              d.testcases = [];
            });
          },
          selectTool: async (toolName: string) => {
            if (!toolName || get().activeTool?.name === toolName) {
              return;
            }
          },
        };
      },
    ),
  ),
);
