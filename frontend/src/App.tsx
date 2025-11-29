import { Link, Switch, Route, useLocation } from 'wouter';
import {
  BrainCircuit,
  ClipboardList,
  Hammer,
  House,
  MessageSquareText,
  Settings,
  UserRound,
} from 'lucide-react';
import * as React from 'react';
import { cn } from '@/lib/utils';
import {
  HomePage,
  ClipboardPage,
  PromptPage,
  LLMPage,
  UserPage,
  SettingPage,
  ToolPage,
} from '@/pages';
import { DarkModeToggle } from '@/components/DarkModeToggle';

type IConComponent = typeof Settings;
type FeatureProps = {
  name: string;
  Icon: IConComponent;
  href: string;
};
const Feature: React.FC<FeatureProps> = ({ name, Icon, href }) => {
  const [location] = useLocation();
  const isActive = href === location;
  return (
    <Link
      key={name}
      href={href}
      className={cn(
        'w-8 h-8 flex items-center justify-center rounded-sm transition-colors',
        isActive
          ? 'bg-primary text-primary-foreground border border-border'
          : 'text-muted-foreground hover:text-foreground hover:bg-accent/50',
      )}
    >
      <Icon size="20" />
    </Link>
  );
};

const features = [
  {
    name: 'Home',
    Icon: House,
    href: '/',
  },
  {
    name: 'Tools',
    Icon: Hammer,
    href: '/tools',
  },
  {
    name: 'Clipboard History',
    Icon: ClipboardList,
    href: '/clipboard',
  },
  {
    name: 'Prompt',
    Icon: MessageSquareText,
    href: '/prompt',
  },
  {
    name: 'LLM',
    Icon: BrainCircuit,
    href: '/llm',
  },
].map((feature) => <Feature key={feature.name} {...feature} />);

const settings = [
  {
    name: 'user',
    Icon: UserRound,
    href: '/user',
  },
  { name: 'setting', Icon: Settings, href: '/setting' },
].map((setting) => <Feature key={setting.name} {...setting} />);

function App() {
  return (
    <div className="w-full h-full flex flex-col bg-background text-foreground">
      <div className="absolute top-4 right-4">
        <DarkModeToggle />
      </div>
      <div className="flex flex-1 min-h-0">
        <div className="w-10 flex flex-col justify-between border-r border-border bg-secondary/30 dark:bg-secondary/10 z-20">
          <div className="text-center flex flex-col items-center pt-2 gap-2">{features}</div>
          <div className="text-center flex flex-col items-center gap-2">{settings}</div>
        </div>
        <div className="flex-1 min-h-0">
          <Switch>
            <Route path="/" component={HomePage} />
            <Route path="/tools" component={ToolPage} />
            <Route path="/clipboard" component={ClipboardPage} />
            <Route path="/prompt" component={PromptPage} />
            <Route path="/llm" component={LLMPage} />
            <Route path="/user" component={UserPage} />
            <Route path="/setting" component={SettingPage} />
          </Switch>
        </div>
      </div>
    </div>
  );
}

export default App;
