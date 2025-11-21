import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import './App.css';

// ============================================================================
// TYPES
// ============================================================================

interface Employee {
  id: string;
  email: string;
  full_name: string;
  role: string;
  department: string;
  hire_date: string;
  status: 'active' | 'inactive';
}

interface Project {
  id: string;
  name: string;
  description: string;
  start_date: string;
  end_date: string;
  status: 'planning' | 'active' | 'completed' | 'on_hold';
  budget: number;
  assigned_employees?: string[];
}

interface Task {
  id: string;
  title: string;
  description: string;
  project_id: string;
  assigned_to: string;
  status: 'todo' | 'in_progress' | 'completed';
  priority: 'low' | 'medium' | 'high';
  due_date: string;
}

// ============================================================================
// MOCK DATA
// ============================================================================

const mockEmployees: Employee[] = [
  {
    id: '1',
    email: 'john.doe@company.com',
    full_name: 'John Doe',
    role: 'Software Engineer',
    department: 'Engineering',
    hire_date: '2023-01-15',
    status: 'active'
  },
  {
    id: '2',
    email: 'jane.smith@company.com',
    full_name: 'Jane Smith',
    role: 'Product Manager',
    department: 'Product',
    hire_date: '2022-06-20',
    status: 'active'
  },
  {
    id: '3',
    email: 'bob.johnson@company.com',
    full_name: 'Bob Johnson',
    role: 'Designer',
    department: 'Design',
    hire_date: '2023-03-10',
    status: 'inactive'
  }
];

const mockProjects: Project[] = [
  {
    id: '1',
    name: 'Website Redesign',
    description: 'Complete overhaul of company website',
    start_date: '2024-01-01',
    end_date: '2024-06-30',
    status: 'active',
    budget: 50000,
    assigned_employees: ['1', '3']
  },
  {
    id: '2',
    name: 'Mobile App Development',
    description: 'Build iOS and Android apps',
    start_date: '2024-03-15',
    end_date: '2024-12-31',
    status: 'planning',
    budget: 120000,
    assigned_employees: ['1', '2']
  },
  {
    id: '3',
    name: 'Data Migration',
    description: 'Migrate legacy data to new system',
    start_date: '2023-06-01',
    end_date: '2023-12-31',
    status: 'completed',
    budget: 30000,
    assigned_employees: ['1']
  }
];

const mockTasks: Task[] = [
  {
    id: '1',
    title: 'Design homepage mockup',
    description: 'Create initial design concepts for new homepage',
    project_id: '1',
    assigned_to: '3',
    status: 'completed',
    priority: 'high',
    due_date: '2024-02-15'
  },
  {
    id: '2',
    title: 'Setup React project',
    description: 'Initialize frontend application with Vite',
    project_id: '1',
    assigned_to: '1',
    status: 'completed',
    priority: 'high',
    due_date: '2024-02-20'
  },
  {
    id: '3',
    title: 'Implement navigation',
    description: 'Build responsive navigation component',
    project_id: '1',
    assigned_to: '1',
    status: 'in_progress',
    priority: 'medium',
    due_date: '2024-03-01'
  },
  {
    id: '4',
    title: 'API research',
    description: 'Research best practices for mobile API design',
    project_id: '2',
    assigned_to: '2',
    status: 'todo',
    priority: 'medium',
    due_date: '2024-04-01'
  }
];

// ============================================================================
// COMPONENTS
// ============================================================================

function Login() {
  return (
    <div className="page">
      <h1>Login</h1>
      <p>Login form will go here (Phase 1B)</p>
    </div>
  );
}

function Dashboard() {
  return (
    <div className="page">
      <h1>Dashboard</h1>
      <div className="stats">
        <h2>Quick Stats</h2>
        <p>Total Employees: {mockEmployees.length}</p>
        <p>Active Projects: {mockProjects.filter(p => p.status === 'active').length}</p>
        <p>Pending Tasks: {mockTasks.filter(t => t.status !== 'completed').length}</p>
      </div>
    </div>
  );
}

function Employees() {
  return (
    <div className="page">
      <h1>Employees</h1>
      <button>Add New Employee</button>

      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Email</th>
            <th>Role</th>
            <th>Department</th>
            <th>Hire Date</th>
            <th>Status</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {mockEmployees.map((employee) => (
            <tr key={employee.id}>
              <td>{employee.full_name}</td>
              <td>{employee.email}</td>
              <td>{employee.role}</td>
              <td>{employee.department}</td>
              <td>{employee.hire_date}</td>
              <td>{employee.status}</td>
              <td>
                <button>Edit</button>
                <button>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function Projects() {
  return (
    <div className="page">
      <h1>Projects</h1>
      <button>Add New Project</button>

      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Description</th>
            <th>Start Date</th>
            <th>End Date</th>
            <th>Status</th>
            <th>Budget</th>
            <th>Team Size</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {mockProjects.map((project) => (
            <tr key={project.id}>
              <td>{project.name}</td>
              <td>{project.description}</td>
              <td>{project.start_date}</td>
              <td>{project.end_date}</td>
              <td>
                <span className={`status-badge status-${project.status}`}>
                  {project.status}
                </span>
              </td>
              <td>${project.budget.toLocaleString()}</td>
              <td>{project.assigned_employees?.length || 0} employees</td>
              <td>
                <button>Edit</button>
                <button>View Team</button>
                <button>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

function Tasks() {
  return (
    <div className="page">
      <h1>Tasks</h1>
      <button>Add New Task</button>

      <table>
        <thead>
          <tr>
            <th>Title</th>
            <th>Project</th>
            <th>Assigned To</th>
            <th>Status</th>
            <th>Priority</th>
            <th>Due Date</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {mockTasks.map((task) => {
            const project = mockProjects.find(p => p.id === task.project_id);
            const employee = mockEmployees.find(e => e.id === task.assigned_to);

            return (
              <tr key={task.id}>
                <td>{task.title}</td>
                <td>{project?.name || 'Unknown'}</td>
                <td>{employee?.full_name || 'Unassigned'}</td>
                <td>
                  <span className={`status-badge status-${task.status}`}>
                    {task.status}
                  </span>
                </td>
                <td>
                  <span className={`priority-badge priority-${task.priority}`}>
                    {task.priority}
                  </span>
                </td>
                <td>{task.due_date}</td>
                <td>
                  <button>Edit</button>
                  <button>Delete</button>
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
  );
}

// ============================================================================
// APP
// ============================================================================

function App() {
  return (
    <Router>
      <nav>
        <ul>
          <li><Link to="/">Login</Link></li>
          <li><Link to="/dashboard">Dashboard</Link></li>
          <li><Link to="/employees">Employees</Link></li>
          <li><Link to="/projects">Projects</Link></li>
          <li><Link to="/tasks">Tasks</Link></li>
        </ul>
      </nav>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/employees" element={<Employees />} />
        <Route path="/projects" element={<Projects />} />
        <Route path="/tasks" element={<Tasks />} />
      </Routes>
    </Router>
  );
}

export default App;
