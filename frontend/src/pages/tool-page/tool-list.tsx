import type { HTMLAttributes } from 'react';
import * as React from 'react';
import { Plus, SquareTerminal } from 'lucide-react';
import { Button } from '@/components/ui/button';
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarTrigger,
} from '@/components/ui/sidebar';
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from '@/components/ui/context-menu';
import { useToolPageStore, useShallow } from './store';
import { cn } from '@/lib/utils';

type ToolListProps = HTMLAttributes<HTMLDivElement> & {};
export function ToolList({}: ToolListProps) {
  const { tools, toolsDir, selectTool } = useToolPageStore(
    useShallow((s) => ({
      tools: s.tools,
      toolsDir: s.toolsDir,
      selectTool: s.selectTool,
    })),
  );
  const activeToolName = useToolPageStore((s) => s.activeTool?.name || '');

  const list = React.useMemo(() => {
    if (!Array.isArray(tools) || tools.length === 0) {
      return [];
    }
    return tools.map((tool) => ({
      title: tool.name,
      isActive: tool.name === activeToolName,
      onClick: () => {
        selectTool(tool.name);
      },
      icon: SquareTerminal,
      description: tool.description,
    }));
  }, [tools, activeToolName]);

  return (
    <Sidebar variant="inset" className="left-10 p-0 bg-secondary/30 dark:bg-secondary/10">
      <SidebarContent className="bg-background">
        <SidebarGroup className="p-0">
          <SidebarGroupLabel className="flex justify-between bg-muted dark:bg-muted border-b border-solid border-border rounded-none pl-4"></SidebarGroupLabel>
          <SidebarGroupContent className="p-0 pt-2">
            <SidebarMenu>
              {list.map((item) => (
                <SidebarMenuItem key={item.title}>
                  <SidebarMenuButton
                    onClick={item.onClick}
                    aria-current={item.isActive ? 'page' : undefined}
                    className="w-full"
                  >
                    <div
                      className={cn(
                        'flex items-center cursor-pointer gap-2 px-2 py-1 rounded-md transition-colors w-full',
                        item.isActive
                          ? 'bg-primary text-primary-foreground'
                          : 'text-muted-foreground hover:bg-accent/50 hover:text-foreground',
                      )}
                    >
                      <item.icon
                        className={cn(
                          item.isActive
                            ? 'text-primary-foreground'
                            : 'text-muted-foreground hover:text-foreground',
                        )}
                        strokeWidth={item.isActive ? 2 : 1}
                      />
                      <span className="truncate select-none">{item.title}</span>
                    </div>
                  </SidebarMenuButton>
                </SidebarMenuItem>
              ))}
            </SidebarMenu>
          </SidebarGroupContent>
        </SidebarGroup>
      </SidebarContent>
    </Sidebar>
  );
}
