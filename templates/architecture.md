# Architecture du système

## Vue d'ensemble

Cette documentation décrit l'architecture technique du projet [nom du projet], ses composants principaux, leurs interactions, et les décisions architecturales prises pour répondre aux exigences fonctionnelles et non-fonctionnelles.

### Objectifs architecturaux

- **Scalabilité** : Capacité à gérer une charge croissante d'utilisateurs
- **Maintenabilité** : Code structuré et facilement modifiable
- **Performance** : Temps de réponse optimaux pour une bonne UX
- **Sécurité** : Protection des données et accès sécurisés
- **Disponibilité** : Haute disponibilité avec un minimum de downtime
- **Évolutivité** : Facilité d'ajout de nouvelles fonctionnalités

## Architecture générale

### Diagramme de haut niveau

```
┌─────────────────────────────────────────────────────┐
│                    UTILISATEURS                     │
│              (Web, Mobile, API)                     │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                LOAD BALANCER                        │
│              (Nginx/HAProxy)                        │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                   CDN/CACHE                         │
│              (CloudFlare/Redis)                     │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                 FRONTEND LAYER                      │
│          (React/Vue/Angular + Static Assets)        │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                 API GATEWAY                         │
│            (Authentication, Routing)                │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│               BACKEND SERVICES                      │
│         ┌─────────┬─────────┬─────────┐            │
│         │   API   │ Worker  │ Cron    │            │
│         │ Service │ Service │ Service │            │
│         └─────────┴─────────┴─────────┘            │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                DATA LAYER                           │
│    ┌────────────┬────────────┬────────────┐        │
│    │  Database  │   Cache    │  File      │        │
│    │ (Primary)  │  (Redis)   │ Storage    │        │
│    └────────────┴────────────┴────────────┘        │
└─────────────────────────────────────────────────────┘
```

### Patterns architecturaux utilisés

- **Architecture en couches** (Layered Architecture)
- **Microservices** ou **Service-Oriented Architecture (SOA)**
- **Model-View-Controller (MVC)**
- **Repository Pattern** pour l'accès aux données
- **Observer Pattern** pour les événements
- **Circuit Breaker** pour la résilience
- **CQRS** (Command Query Responsibility Segregation) [si applicable]

## Architecture frontend

### Stack technologique

- **Framework** : [React/Vue.js/Angular] v[version]
- **State Management** : [Redux/Vuex/NgRx/Zustand]
- **Router** : [React Router/Vue Router/Angular Router]
- **Styling** : [Tailwind CSS/Styled Components/SCSS]
- **Build Tool** : [Vite/Webpack/Parcel]
- **Testing** : [Jest/Vitest + Testing Library]

### Structure des composants

```
src/
├── components/           # Composants réutilisables
│   ├── ui/              # Composants UI de base (Button, Input)
│   ├── forms/           # Composants de formulaires
│   ├── layout/          # Composants de mise en page
│   └── business/        # Composants métier spécifiques
├── pages/               # Pages de l'application
│   ├── auth/           # Pages d'authentification
│   ├── dashboard/      # Pages du tableau de bord
│   └── settings/       # Pages de paramètres
├── hooks/               # Custom hooks React
├── services/            # Services API et logique métier
├── stores/              # Gestion d'état (Redux/Zustand)
├── utils/               # Fonctions utilitaires
├── types/               # Définitions TypeScript
└── assets/              # Assets statiques
```

### Architecture des composants

```typescript
// Exemple d'architecture de composant
interface ComponentProps {
  // Props typées
}

interface ComponentState {
  // State local
}

// Composant fonctionnel avec hooks
const MyComponent: React.FC<ComponentProps> = ({ prop1, prop2 }) => {
  // Hooks pour l'état local
  const [state, setState] = useState<ComponentState>();

  // Hooks pour l'état global
  const globalState = useSelector(selectSomeData);
  const dispatch = useDispatch();

  // Hooks personnalisés
  const { data, loading, error } = useApiData();

  // Effects
  useEffect(() => {
    // Side effects
  }, []);

  // Event handlers
  const handleAction = useCallback(() => {
    // Action logic
  }, []);

  // Render
  return (
    <div>
      {/* JSX */}
    </div>

### Gestion d'état

#### État global (Redux/Zustand)
```typescript
// Store structure
interface AppState {
  auth: AuthState;
  user: UserState;
  products: ProductsState;
  ui: UIState;
}

// Slices/Modules
const authSlice = createSlice({
  name: 'auth',
  initialState: {
    user: null,
    token: null,
    isAuthenticated: false,
    loading: false
  },
  reducers: {
    loginStart: (state) => { state.loading = true; },
    loginSuccess: (state, action) => {
      state.user = action.payload.user;
      state.token = action.payload.token;
      state.isAuthenticated = true;
      state.loading = false;
    },
    loginFailure: (state) => { state.loading = false; },
    logout: (state) => {
      state.user = null;
      state.token = null;
      state.isAuthenticated = false;
    }
  }
});
```

#### État local optimisé
```typescript
// Custom hooks pour la logique réutilisable
function useFormValidation<T>(initialValues: T, validationRules: ValidationRules<T>) {
  const [values, setValues] = useState(initialValues);
  const [errors, setErrors] = useState<ValidationErrors<T>>({});
  const [touched, setTouched] = useState<Record<keyof T, boolean>>({} as any);

  const validate = useCallback((fieldName?: keyof T) => {
    // Validation logic
  }, [values, validationRules]);

  const handleChange = useCallback((name: keyof T, value: any) => {
    setValues(prev => ({ ...prev, [name]: value }));
    if (touched[name]) validate(name);
  }, [touched, validate]);

  return { values, errors, touched, handleChange, validate };
}
```

### Routing et navigation

```typescript
// Structure de routing
const AppRouter: React.FC = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<PublicLayout />}>
          <Route index element={<HomePage />} />
          <Route path="auth/*" element={<AuthRoutes />} />
        </Route>

        <Route path="/app" element={<PrivateLayout />}>
          <Route index element={<Dashboard />} />
          <Route path="users/*" element={<UserRoutes />} />
          <Route path="products/*" element={<ProductRoutes />} />
          <Route path="settings/*" element={<SettingsRoutes />} />
        </Route>

        <Route path="*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
};

// Route guards
const PrivateRoute: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const isAuthenticated = useSelector(selectIsAuthenticated);
  const location = useLocation();

  if (!isAuthenticated) {
    return <Navigate to="/auth/login" state={{ from: location }} replace />;
  }

  return <>{children}</>;
};
```

## Architecture backend

### Stack technologique

- **Runtime** : [Node.js/Python/Go/Java] v[version]
- **Framework** : [Express/FastAPI/Gin/Spring Boot]
- **Base de données** : [PostgreSQL/MongoDB/MySQL] v[version]
- **Cache** : [Redis/Memcached] v[version]
- **Message Queue** : [RabbitMQ/Apache Kafka/AWS SQS]
- **Search Engine** : [Elasticsearch/Algolia] [si applicable]

### Architecture en couches

```
├── src/
│   ├── controllers/         # Contrôleurs API (HTTP handlers)
│   ├── services/           # Logique métier
│   ├── repositories/       # Accès aux données
│   ├── models/             # Modèles de données
│   ├── middleware/         # Middleware Express/FastAPI
│   ├── validators/         # Validation des données
│   ├── utils/              # Utilitaires
│   ├── config/             # Configuration
│   ├── database/           # Migrations, seeds
│   ├── tests/              # Tests unitaires et d'intégration
│   └── types/              # Types TypeScript/interfaces
```

### Couche contrôleur

```typescript
// Exemple de contrôleur REST
@Controller('/api/users')
export class UserController {
  constructor(
    private userService: UserService,
    private logger: Logger
  ) {}

  @Get('/')
  @UseMiddleware(authMiddleware, validateQuery)
  async getUsers(
    @Query() query: GetUsersQuery,
    @Req() req: AuthenticatedRequest
  ): Promise<ApiResponse<PaginatedResult<User>>> {
    try {
      const result = await this.userService.getUsers(query, req.user);
      return successResponse(result);
    } catch (error) {
      this.logger.error('Error fetching users', error);
      throw error;
    }
  }

  @Post('/')
  @UseMiddleware(authMiddleware, validateBody)
  async createUser(
    @Body() userData: CreateUserDto,
    @Req() req: AuthenticatedRequest
  ): Promise<ApiResponse<User>> {
    const user = await this.userService.createUser(userData, req.user);
    return successResponse(user, 'User created successfully', 201);
  }
}
```

### Couche service (logique métier)

```typescript
// Service avec injection de dépendances
@Injectable()
export class UserService {
  constructor(
    private userRepository: UserRepository,
    private emailService: EmailService,
    private cacheService: CacheService,
    private eventEmitter: EventEmitter
  ) {}

  async getUsers(query: GetUsersQuery, currentUser: User): Promise<PaginatedResult<User>> {
    // Vérification des permissions
    this.checkPermissions(currentUser, 'read:users');

    // Cache check
    const cacheKey = `users:${JSON.stringify(query)}`;
    const cached = await this.cacheService.get(cacheKey);
    if (cached) return cached;

    // Business logic
    const filters = this.buildFilters(query);
    const result = await this.userRepository.findMany(filters, query.pagination);

    // Cache result
    await this.cacheService.set(cacheKey, result, 300); // 5 minutes

    return result;
  }

  async createUser(userData: CreateUserDto, currentUser: User): Promise<User> {
    // Validation métier
    await this.validateUserCreation(userData);

    // Transaction
    return await this.userRepository.transaction(async (trx) => {
      // Créer l'utilisateur
      const user = await this.userRepository.create(userData, trx);

      // Envoyer email de bienvenue
      await this.emailService.sendWelcomeEmail(user.email);

      // Émettre événement
      this.eventEmitter.emit('user.created', { user, createdBy: currentUser });

      return user;
    });
  }

  private checkPermissions(user: User, permission: string): void {
    if (!user.hasPermission(permission)) {
      throw new ForbiddenError('Insufficient permissions');
    }
  }
}
```

### Couche repository (accès aux données)

```typescript
// Repository avec pattern générique
export abstract class BaseRepository<T> {
  constructor(protected db: Database, protected tableName: string) {}

  async findById(id: string): Promise<T | null> {
    const result = await this.db(this.tableName).where({ id }).first();
    return result ? this.mapToEntity(result) : null;
  }

  async findMany(filters: QueryFilters, pagination: Pagination): Promise<PaginatedResult<T>> {
    const query = this.db(this.tableName);

    // Apply filters
    this.applyFilters(query, filters);

    // Count total
    const total = await query.clone().count('* as count').first();

    // Apply pagination
    const offset = (pagination.page - 1) * pagination.limit;
    const items = await query.limit(pagination.limit).offset(offset);

    return {
      items: items.map(item => this.mapToEntity(item)),
      pagination: {
        page: pagination.page,
        limit: pagination.limit,
        total: total.count,
        pages: Math.ceil(total.count / pagination.limit)
      }
    };
  }

  protected abstract mapToEntity(row: any): T;
}

// Implémentation spécifique
@Injectable()
export class UserRepository extends BaseRepository<User> {
  constructor(db: Database) {
    super(db, 'users');
  }

  async findByEmail(email: string): Promise<User | null> {
    const result = await this.db(this.tableName).where({ email }).first();
    return result ? this.mapToEntity(result) : null;
  }

  protected mapToEntity(row: any): User {
    return new User({
      id: row.id,
      email: row.email,
      firstName: row.first_name,
      lastName: row.last_name,
      role: row.role,
      createdAt: row.created_at,
      updatedAt: row.updated_at
    });
  }
}
```

## Architecture des données

### Modèle de données

#### Entités principales

```typescript
// Domain entities
export class User {
  constructor(
    public readonly id: UserId,
    public email: Email,
    public firstName: string,
    public lastName: string,
    public role: UserRole,
    public status: UserStatus,
    public createdAt: Date,
    public updatedAt: Date
  ) {}

  // Business methods
  hasPermission(permission: string): boolean {
    return this.role.hasPermission(permission);
  }

  updateProfile(data: UpdateProfileData): void {
    // Business logic for profile update
    this.firstName = data.firstName || this.firstName;
    this.lastName = data.lastName || this.lastName;
    this.updatedAt = new Date();
  }

  activate(): void {
    if (this.status === UserStatus.PENDING) {
      this.status = UserStatus.ACTIVE;
      this.updatedAt = new Date();
    }
  }
}

// Value objects
export class Email {
  constructor(private readonly value: string) {
    if (!this.isValid(value)) {
      throw new Error('Invalid email format');
    }
  }

  toString(): string {
    return this.value;
  }

  private isValid(email: string): boolean {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return emailRegex.test(email);
  }
}
```

### Schéma de base de données

#### PostgreSQL Schema

```sql
-- Users table
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'user',
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    email_verified_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_created_at ON users(created_at);
```

#### MongoDB Schema

```javascript
// User collection schema
const userSchema = new mongoose.Schema({
  email: {
    type: String,
    required: true,
    unique: true,
    lowercase: true,
    trim: true
  },
  passwordHash: {
    type: String,
    required: true
  },
  firstName: {
    type: String,
    required: true,
    trim: true
  },
  lastName: {
    type: String,
    required: true,
    trim: true
  },
  role: {
    type: String,
    enum: ['admin', 'user', 'moderator'],
    default: 'user'
  },
  status: {
    type: String,
    enum: ['active', 'pending', 'suspended'],
    default: 'pending'
  },
  preferences: {
    newsletter: { type: Boolean, default: true },
    notifications: { type: Boolean, default: true },
    language: { type: String, default: 'fr' }
  }
}, {
  timestamps: true, // Ajoute createdAt et updatedAt automatiquement
  collection: 'users'
});

// Index pour les performances
userSchema.index({ email: 1 });
userSchema.index({ status: 1, role: 1 });
userSchema.index({ createdAt: -1 });
```

### Stratégie de cache

#### Architecture de cache multi-niveau

```typescript
// Cache service avec fallback
@Injectable()
export class CacheService {
  constructor(
    private redis: Redis,
    private memoryCache: MemoryCache
  ) {}

  async get<T>(key: string): Promise<T | null> {
    // Level 1: Memory cache (très rapide)
    let value = this.memoryCache.get<T>(key);
    if (value) return value;

    // Level 2: Redis cache (rapide)
    const redisValue = await this.redis.get(key);
    if (redisValue) {
      value = JSON.parse(redisValue);
      // Repopuler le cache mémoire
      this.memoryCache.set(key, value, 60); // 1 minute
      return value;
    }

    return null;
  }

  async set<T>(key: string, value: T, ttlSeconds: number): Promise<void> {
    // Stocker dans les deux niveaux
    this.memoryCache.set(key, value, Math.min(ttlSeconds, 300)); // Max 5 min en mémoire
    await this.redis.setex(key, ttlSeconds, JSON.stringify(value));
  }

  async invalidate(pattern: string): Promise<void> {
    // Invalider pattern dans Redis
    const keys = await this.redis.keys(pattern);
    if (keys.length > 0) {
      await this.redis.del(...keys);
    }

    // Invalider pattern en mémoire
    this.memoryCache.invalidatePattern(pattern);
  }
}

// Cache decorators pour simplifier l'usage
export function Cached(ttl: number = 300, keyGenerator?: (args: any[]) => string) {
  return function (target: any, propertyName: string, descriptor: PropertyDescriptor) {
    const method = descriptor.value;

    descriptor.value = async function (...args: any[]) {
      const cacheKey = keyGenerator ? keyGenerator(args) : `${target.constructor.name}:${propertyName}:${JSON.stringify(args)}`;

      // Vérifier le cache
      const cached = await this.cacheService.get(cacheKey);
      if (cached) return cached;

      // Exécuter la méthode
      const result = await method.apply(this, args);

      // Mettre en cache
      await this.cacheService.set(cacheKey, result, ttl);

      return result;
    };
  };
}

// Utilisation
export class UserService {
  @Cached(300, (args) => `user:${args[0]}`)
  async getUserById(id: string): Promise<User> {
    return await this.userRepository.findById(id);
  }
}
```

## Architecture de sécurité

### Authentification et autorisation

```typescript
// JWT Service
@Injectable()
export class AuthService {
  constructor(
    private userService: UserService,
    private jwtService: JwtService,
    private bcryptService: BcryptService
  ) {}

  async authenticate(email: string, password: string): Promise<AuthResult> {
    // Vérifier l'utilisateur
    const user = await this.userService.findByEmail(email);
    if (!user || !await this.bcryptService.compare(password, user.passwordHash)) {
      throw new UnauthorizedError('Invalid credentials');
    }

    // Vérifier le statut
    if (user.status !== UserStatus.ACTIVE) {
      throw new ForbiddenError('Account not active');
    }

    // Générer les tokens
    const accessToken = this.generateAccessToken(user);
    const refreshToken = this.generateRefreshToken(user);

    // Enregistrer le refresh token
    await this.storeRefreshToken(user.id, refreshToken);

    return {
      accessToken,
      refreshToken,
      user: user.toPublicData(),
      expiresIn: 3600 // 1 heure
    };
  }

  private generateAccessToken(user: User): string {
    const payload = {
      sub: user.id,
      email: user.email,
      role: user.role,
      permissions: user.getPermissions()
    };

    return this.jwtService.sign(payload, {
      expiresIn: '1h',
      issuer: 'api.example.com',
      audience: 'app.example.com'
    });
  }
}

// Authorization middleware
export function authorize(permissions: string[]) {
  return (req: AuthenticatedRequest, res: Response, next: NextFunction) => {
    const user = req.user;

    if (!user) {
      return res.status(401).json({ message: 'Authentication required' });
    }

    const hasPermission = permissions.some(permission =>
      user.hasPermission(permission)
    );

    if (!hasPermission) {
      return res.status(403).json({ message: 'Insufficient permissions' });
    }

    next();
  };
}
```

### Sécurité des données

```typescript
// Data encryption service
@Injectable()
export class EncryptionService {
  private readonly algorithm = 'aes-256-gcm';
  private readonly key: Buffer;

  constructor() {
    this.key = Buffer.from(process.env.ENCRYPTION_KEY!, 'hex');
  }

  encrypt(text: string): EncryptedData {
    const iv = crypto.randomBytes(16);
    const cipher = crypto.createCipher(this.algorithm, this.key);
    cipher.setAAD(Buffer.from('additional-data'));

    let encrypted = cipher.update(text, 'utf8', 'hex');
    encrypted += cipher.final('hex');

    const authTag = cipher.getAuthTag();

    return {
      encrypted,
      iv: iv.toString('hex'),
      authTag: authTag.toString('hex')
    };
  }

  decrypt(data: EncryptedData): string {
    const decipher = crypto.createDecipher(this.algorithm, this.key);
    decipher.setAAD(Buffer.from('additional-data'));
    decipher.setAuthTag(Buffer.from(data.authTag, 'hex'));

    let decrypted = decipher.update(data.encrypted, 'hex', 'utf8');
    decrypted += decipher.final('utf8');

    return decrypted;
  }
}

// Sensitive data handling
export class SensitiveField {
  constructor(
    private encryptionService: EncryptionService,
    private value: string
  ) {}

  encrypt(): EncryptedData {
    return this.encryptionService.encrypt(this.value);
  }

  static decrypt(encryptionService: EncryptionService, data: EncryptedData): string {
    return encryptionService.decrypt(data);
  }
}
```

## Performance et optimisation

### Stratégies d'optimisation

#### Database Query Optimization

```typescript
// Query builder avec optimisations
export class OptimizedQueryBuilder {
  constructor(private db: Database) {}

  // Eager loading pour éviter N+1
  async findUsersWithOrders(): Promise<UserWithOrders[]> {
    return await this.db('users')
      .select([
        'users.*',
        this.db.raw('json_agg(orders.*) as orders')
      ])
      .leftJoin('orders', 'users.id', 'orders.user_id')
      .groupBy('users.id');
  }

  // Pagination optimisée avec cursor
  async findWithCursor(cursor?: string, limit: number = 20): Promise<CursorResult<User>> {
    const query = this.db('users').orderBy('created_at', 'desc').limit(limit + 1);

    if (cursor) {
      const cursorData = Buffer.from(cursor, 'base64').toString();
      const [createdAt, id] = cursorData.split('|');
      query.where(function() {
        this.where('created_at', '<', createdAt)
          .orWhere(function() {
            this.where('created_at', '=', createdAt)
              .andWhere('id', '<', id);
          });
      });
    }

    const results = await query;
    const hasMore = results.length > limit;
    const items = hasMore ? results.slice(0, -1) : results;

    const nextCursor = hasMore ?
      Buffer.from(`${items[items.length - 1].created_at}|${items[items.length - 1].id}`).toString('base64') :
      null;

    return { items, nextCursor, hasMore };
  }
}
```

#### Caching Strategy

```typescript
// Cache warming et invalidation intelligente
@Injectable()
export class CacheWarmingService {
  constructor(
    private cacheService: CacheService,
    private userService: UserService
  ) {}

  @Cron('0 */6 * * *') // Toutes les 6 heures
  async warmFrequentlyAccessedData(): Promise<void> {
    // Warming des données fréquemment accédées
    const popularUsers = await this.userService.getMostActiveUsers(100);

    await Promise.all(
      popularUsers.map(user =>
        this.cacheService.set(`user:${user.id}`, user, 21600) // 6 heures
      )
    );
  }

  // Invalidation en cascade
  async invalidateUserCache(userId: string): Promise<void> {
    await Promise.all([
      this.cacheService.invalidate(`user:${userId}`),
      this.cacheService.invalidate(`user:${userId}:*`),
      this.cacheService.invalidate(`user-orders:${userId}:*`),
      this.cacheService.invalidate(`user-permissions:${userId}`)
    ]);
  }
}
```

### Monitoring et observabilité

```typescript
// Metrics collection
@Injectable()
export class MetricsService {
  private registry = new Registry();

  // Compteurs
  private httpRequestsTotal = new Counter({
    name: 'http_requests_total',
    help: 'Total number of HTTP requests',
    labelNames: ['method', 'route', 'status'],
    registers: [this.registry]
  });

  // Histogrammes pour la latence
  private httpRequestDuration = new Histogram({
    name: 'http_request_duration_seconds',
    help: 'HTTP request duration in seconds',
    labelNames: ['method', 'route'],
    buckets: [0.1, 0.5, 1, 2, 5],
    registers: [this.registry]
  });

  // Jauges pour les métriques temps réel
  private activeConnections = new Gauge({
    name: 'active_connections',
    help: 'Number of active connections',
    registers: [this.registry]
  });

  recordHttpRequest(method: string, route: string, status: number, duration: number): void {
    this.httpRequestsTotal.inc({ method, route, status: status.toString() });
    this.httpRequestDuration.observe({ method, route }, duration);
  }

  setActiveConnections(count: number): void {
    this.activeConnections.set(count);
  }

  getMetrics(): string {
    return this.registry.metrics();
  }
}

// Middleware de monitoring
export function monitoringMiddleware(metricsService: MetricsService) {
  return (req: Request, res: Response, next: NextFunction) => {
    const start = Date.now();

    res.on('finish', () => {
      const duration = (Date.now() - start) / 1000;
      metricsService.recordHttpRequest(
        req.method,
        req.route?.path || req.path,
        res.statusCode,
        duration
      );
    });

    next();
  };
}
```

## Déploiement et infrastructure

### Architecture cloud

```yaml
# Docker Compose pour développement local
version: '3.8'
services:
  # Frontend
  frontend:
    build: ./frontend
    ports:
      - "3000:3000"
    environment:
      - REACT_APP_API_URL=http://localhost:8000
    depends_on:
      - backend

  # Backend API
  backend:
    build: ./backend
    ports:
      - "8000:8000"
    environment:
      - DATABASE_URL=postgresql://user:pass@postgres:5432/myapp
      - REDIS_URL=redis://redis:6379
      - JWT_SECRET=${JWT_SECRET}
    depends_on:
      - postgres
      - redis

  # Base de données
  postgres:
    image: postgres:15
    environment:
      - POSTGRES_DB=myapp
      - POSTGRES_USER=user
      - POSTGRES_PASSWORD=pass
    volumes:
      - postgres_data:/var/lib/postgresql/data
      - ./database/init:/docker-entrypoint-initdb.d

  # Cache
  redis:
    image: redis:7-alpine
    command: redis-server --appendonly yes
    volumes:
      - redis_data:/data

  # Message Queue
  rabbitmq:
    image: rabbitmq:3-management
    environment:
      - RABBITMQ_DEFAULT_USER=admin
      - RABBITMQ_DEFAULT_PASS=password
    ports:
      - "15672:15672" # Management UI

volumes:
  postgres_data:
  redis_data:
```

### Configuration Kubernetes

```yaml
# Kubernetes deployment
apiVersion: apps/v1
kind: Deployment
metadata:
  name: backend-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: backend
  template:
    metadata:
      labels:
        app: backend
    spec:
      containers:
      - name: backend
        image: myapp/backend:latest
        ports:
        - containerPort: 8000
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: database-url
        - name: REDIS_URL
          valueFrom:
            configMapKeyRef:
              name: app-config
              key: redis-url
        resources:
          requests:
            memory: "256Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "500m"
        livenessProbe:
          httpGet:
            path: /health
            port: 8000
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /ready
            port: 8000
          initialDelaySeconds: 5
          periodSeconds: 5

---
apiVersion: v1
kind: Service
metadata:
  name: backend-service
spec:
  selector:
    app: backend
  ports:
  - protocol: TCP
    port: 80
    targetPort: 8000
  type: ClusterIP
```

### CI/CD Pipeline

```yaml
# GitHub Actions workflow
name: Deploy Application

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5

    steps:
    - uses: actions/checkout@v3

    - name: Setup Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'
        cache: 'npm'

    - name: Install dependencies
      run: npm ci

    - name: Run linting
      run: npm run lint

    - name: Run tests
      run: npm run test:coverage
      env:
        DATABASE_URL: postgresql://postgres:postgres@localhost:5432/test

    - name: Upload coverage
      uses: codecov/codecov-action@v3

  build-and-deploy:
    needs: test
    runs-on: ubuntu-latest
    if: github.ref == 'refs/heads/main'

    steps:
    - uses: actions/checkout@v3

    - name: Configure AWS credentials
      uses: aws-actions/configure-aws-credentials@v2
      with:
        aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
        aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
        aws-region: eu-west-1

    - name: Build and push Docker images
      run: |
        # Build frontend
        docker build -t myapp/frontend:${{ github.sha }} ./frontend
        docker tag myapp/frontend:${{ github.sha }} myapp/frontend:latest

        # Build backend
        docker build -t myapp/backend:${{ github.sha }} ./backend
        docker tag myapp/backend:${{ github.sha }} myapp/backend:latest

        # Push to registry
        docker push myapp/frontend:${{ github.sha }}
        docker push myapp/backend:${{ github.sha }}

    - name: Deploy to EKS
      run: |
        aws eks update-kubeconfig --name production-cluster
        kubectl set image deployment/frontend-deployment frontend=myapp/frontend:${{ github.sha }}
        kubectl set image deployment/backend-deployment backend=myapp/backend:${{ github.sha }}
        kubectl rollout status deployment/frontend-deployment
        kubectl rollout status deployment/backend-deployment
```

## Patterns et bonnes pratiques

### Design Patterns utilisés

#### Repository Pattern
```typescript
// Interface générique
interface Repository<T, ID> {
  findById(id: ID): Promise<T | null>;
  findAll(): Promise<T[]>;
  save(entity: T): Promise<T>;
  delete(id: ID): Promise<void>;
  exists(id: ID): Promise<boolean>;
}

// Implémentation abstraite
abstract class BaseRepository<T, ID> implements Repository<T, ID> {
  constructor(protected db: Database, protected tableName: string) {}

  async findById(id: ID): Promise<T | null> {
    const row = await this.db(this.tableName).where({ id }).first();
    return row ? this.mapToEntity(row) : null;
  }

  async save(entity: T): Promise<T> {
    const data = this.mapToRow(entity);
    const [row] = await this.db(this.tableName)
      .insert(data)
      .onConflict('id')
      .merge()
      .returning('*');
    return this.mapToEntity(row);
  }

  protected abstract mapToEntity(row: any): T;
  protected abstract mapToRow(entity: T): any;
}
```

#### Factory Pattern
```typescript
// Factory pour créer des services selon le contexte
interface PaymentService {
  processPayment(amount: number, currency: string): Promise<PaymentResult>;
}

class PaymentServiceFactory {
  static create(provider: PaymentProvider): PaymentService {
    switch (provider) {
      case PaymentProvider.STRIPE:
        return new StripePaymentService();
      case PaymentProvider.PAYPAL:
        return new PayPalPaymentService();
      case PaymentProvider.BANK_TRANSFER:
        return new BankTransferService();
      default:
        throw new Error(`Unsupported payment provider: ${provider}`);
    }
  }
}

// Utilisation
const paymentService = PaymentServiceFactory.create(user.preferredPaymentProvider);
const result = await paymentService.processPayment(100, 'EUR');
```

#### Observer Pattern pour les événements
```typescript
// Event system
interface DomainEvent {
  eventType: string;
  aggregateId: string;
  timestamp: Date;
  payload: any;
}

class EventEmitter {
  private handlers = new Map<string, Function[]>();

  on(eventType: string, handler: Function): void {
    if (!this.handlers.has(eventType)) {
      this.handlers.set(eventType, []);
    }
    this.handlers.get(eventType)!.push(handler);
  }

  async emit(event: DomainEvent): Promise<void> {
    const handlers = this.handlers.get(event.eventType) || [];
    await Promise.all(handlers.map(handler => handler(event)));
  }
}

// Event handlers
@EventHandler('user.created')
export class SendWelcomeEmailHandler {
  constructor(private emailService: EmailService) {}

  async handle(event: DomainEvent): Promise<void> {
    const { user } = event.payload;
    await this.emailService.sendWelcomeEmail(user.email);
  }
}

@EventHandler('order.completed')
export class UpdateInventoryHandler {
  constructor(private inventoryService: InventoryService) {}

  async handle(event: DomainEvent): Promise<void> {
    const { order } = event.payload;
    await this.inventoryService.updateStockLevels(order.items);
  }
}
```

#### Command Pattern pour CQRS
```typescript
// Command base
interface Command {
  readonly commandType: string;
}

interface CommandHandler<T extends Command> {
  handle(command: T): Promise<void>;
}

// Commande spécifique
export class CreateUserCommand implements Command {
  readonly commandType = 'CREATE_USER';

  constructor(
    public readonly email: string,
    public readonly firstName: string,
    public readonly lastName: string,
    public readonly password: string
  ) {}
}

// Handler de commande
@CommandHandler(CreateUserCommand)
export class CreateUserCommandHandler implements CommandHandler<CreateUserCommand> {
  constructor(
    private userRepository: UserRepository,
    private passwordService: PasswordService,
    private eventBus: EventBus
  ) {}

  async handle(command: CreateUserCommand): Promise<void> {
    // Validation
    await this.validateCommand(command);

    // Business logic
    const hashedPassword = await this.passwordService.hash(command.password);
    const user = new User({
      email: command.email,
      firstName: command.firstName,
      lastName: command.lastName,
      passwordHash: hashedPassword
    });

    // Persistence
    await this.userRepository.save(user);

    // Event
    await this.eventBus.publish(new UserCreatedEvent(user));
  }

  private async validateCommand(command: CreateUserCommand): Promise<void> {
    const existingUser = await this.userRepository.findByEmail(command.email);
    if (existingUser) {
      throw new ConflictError('User with this email already exists');
    }
  }
}
```

### Principes SOLID

#### Single Responsibility Principle
```typescript
// ❌ Violation du SRP
class UserService {
  async createUser(userData: CreateUserDto): Promise<User> {
    // Validation
    if (!userData.email || !userData.password) {
      throw new Error('Invalid data');
    }

    // Business logic
    const user = new User(userData);
    await this.userRepository.save(user);

    // Email sending
    const emailTemplate = `Welcome ${user.firstName}!`;
    const transport = nodemailer.createTransporter({...});
    await transport.sendMail({
      to: user.email,
      subject: 'Welcome',
      html: emailTemplate
    });

    return user;
  }
}

// ✅ Respect du SRP
class UserService {
  constructor(
    private userRepository: UserRepository,
    private userValidator: UserValidator,
    private emailService: EmailService
  ) {}

  async createUser(userData: CreateUserDto): Promise<User> {
    await this.userValidator.validate(userData);

    const user = new User(userData);
    await this.userRepository.save(user);

    await this.emailService.sendWelcomeEmail(user);

    return user;
  }
}

class UserValidator {
  async validate(userData: CreateUserDto): Promise<void> {
    if (!userData.email || !userData.password) {
      throw new ValidationError('Email and password are required');
    }
    // Autres validations...
  }
}

class EmailService {
  async sendWelcomeEmail(user: User): Promise<void> {
    const template = await this.templateService.getWelcomeTemplate();
    await this.mailTransport.send({
      to: user.email,
      template: template.render({ user })
    });
  }
}
```

#### Dependency Inversion Principle
```typescript
// Interface pour l'abstraction
interface NotificationService {
  send(recipient: string, message: string): Promise<void>;
}

// Implémentations concrètes
class EmailNotificationService implements NotificationService {
  async send(recipient: string, message: string): Promise<void> {
    // Logique d'envoi d'email
  }
}

class SMSNotificationService implements NotificationService {
  async send(recipient: string, message: string): Promise<void> {
    // Logique d'envoi de SMS
  }
}

// Service de haut niveau qui dépend de l'abstraction
class OrderService {
  constructor(
    private orderRepository: OrderRepository,
    private notificationService: NotificationService // Abstraction, pas implémentation
  ) {}

  async processOrder(order: Order): Promise<void> {
    await this.orderRepository.save(order);
    await this.notificationService.send(
      order.customerEmail,
      `Your order ${order.id} has been processed`
    );
  }
}
```

### Error Handling Strategy

```typescript
// Hiérarchie d'erreurs personnalisées
export abstract class AppError extends Error {
  abstract statusCode: number;
  abstract isOperational: boolean;

  constructor(message: string, public context?: any) {
    super(message);
    Object.setPrototypeOf(this, AppError.prototype);
  }
}

export class ValidationError extends AppError {
  statusCode = 400;
  isOperational = true;

  constructor(message: string, public field?: string) {
    super(message);
  }
}

export class NotFoundError extends AppError {
  statusCode = 404;
  isOperational = true;
}

export class UnauthorizedError extends AppError {
  statusCode = 401;
  isOperational = true;
}

// Global error handler
export class ErrorHandler {
  constructor(private logger: Logger) {}

  handleError(error: Error): void {
    if (this.isOperationalError(error)) {
      this.logger.warn('Operational error', { error: error.message, stack: error.stack });
    } else {
      this.logger.error('Unexpected error', { error: error.message, stack: error.stack });
      // Potentially restart the process or alert monitoring
    }
  }

  private isOperationalError(error: Error): boolean {
    return error instanceof AppError && error.isOperational;
  }
}

// Express middleware
export function errorMiddleware(
  error: Error,
  req: Request,
  res: Response,
  next: NextFunction
): void {
  const errorHandler = new ErrorHandler(logger);
  errorHandler.handleError(error);

  if (error instanceof AppError) {
    res.status(error.statusCode).json({
      success: false,
      message: error.message,
      ...(process.env.NODE_ENV === 'development' && { stack: error.stack })
    });
  } else {
    res.status(500).json({
      success: false,
      message: 'Internal server error'
    });
  }
}
```

## Sécurité avancée

### Input Validation et Sanitization

```typescript
// Validation avec Joi/Zod
import { z } from 'zod';

const CreateUserSchema = z.object({
  email: z.string().email().max(255),
  firstName: z.string().min(1).max(100).regex(/^[a-zA-ZÀ-ÿ\s]+$/),
  lastName: z.string().min(1).max(100).regex(/^[a-zA-ZÀ-ÿ\s]+$/),
  password: z.string()
    .min(8)
    .regex(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]/)
});

// Middleware de validation
export function validateBody<T>(schema: z.ZodSchema<T>) {
  return (req: Request, res: Response, next: NextFunction) => {
    try {
      const validated = schema.parse(req.body);
      req.body = validated;
      next();
    } catch (error) {
      if (error instanceof z.ZodError) {
        const validationErrors = error.errors.map(err => ({
          field: err.path.join('.'),
          message: err.message,
          code: err.code
        }));

        res.status(400).json({
          success: false,
          message: 'Validation failed',
          errors: validationErrors
        });
      } else {
        next(error);
      }
    }
  };
}

// Sanitization
export class Sanitizer {
  static sanitizeHtml(input: string): string {
    return DOMPurify.sanitize(input, {
      ALLOWED_TAGS: ['b', 'i', 'em', 'strong', 'p', 'br'],
      ALLOWED_ATTR: []
    });
  }

  static sanitizeSQL(input: string): string {
    // Utiliser des requêtes préparées plutôt que de la sanitization manuelle
    return input.replace(/[';\\]/g, '');
  }

  static normalizeEmail(email: string): string {
    return email.toLowerCase().trim();
  }
}
```

### Rate Limiting avancé

```typescript
// Rate limiter avec Redis
export class AdvancedRateLimiter {
  constructor(private redis: Redis) {}

  async checkRateLimit(
    identifier: string,
    windowMs: number,
    maxRequests: number,
    strategy: 'fixed' | 'sliding' = 'sliding'
  ): Promise<RateLimitResult> {
    if (strategy === 'sliding') {
      return this.slidingWindowRateLimit(identifier, windowMs, maxRequests);
    } else {
      return this.fixedWindowRateLimit(identifier, windowMs, maxRequests);
    }
  }

  private async slidingWindowRateLimit(
    identifier: string,
    windowMs: number,
    maxRequests: number
  ): Promise<RateLimitResult> {
    const key = `rate_limit:${identifier}`;
    const now = Date.now();
    const windowStart = now - windowMs;

    // Remove old entries
    await this.redis.zremrangebyscore(key, '-inf', windowStart);

    // Count current requests
    const current = await this.redis.zcard(key);

    if (current >= maxRequests) {
      const oldest = await this.redis.zrange(key, 0, 0, 'WITHSCORES');
      const resetTime = oldest.length > 0 ? parseInt(oldest[1]) + windowMs : now + windowMs;

      return {
        allowed: false,
        remaining: 0,
        resetTime,
        retryAfter: Math.ceil((resetTime - now) / 1000)
      };
    }

    // Add current request
    await this.redis.zadd(key, now, `${now}-${Math.random()}`);
    await this.redis.expire(key, Math.ceil(windowMs / 1000));

    return {
      allowed: true,
      remaining: maxRequests - current - 1,
      resetTime: now + windowMs,
      retryAfter: 0
    };
  }
}

// Middleware avec différents niveaux
export function createRateLimitMiddleware(options: RateLimitOptions) {
  const limiter = new AdvancedRateLimiter(redis);

  return async (req: Request, res: Response, next: NextFunction) => {
    const identifier = req.ip + (req.user?.id || '');

    // Différentes limites selon l'endpoint
    let limits: RateLimitConfig;
    if (req.path.startsWith('/api/auth')) {
      limits = { windowMs: 15 * 60 * 1000, maxRequests: 5 }; // Auth: 5/15min
    } else if (req.user?.role === 'premium') {
      limits = { windowMs: 60 * 60 * 1000, maxRequests: 10000 }; // Premium: 10k/hour
    } else {
      limits = { windowMs: 60 * 60 * 1000, maxRequests: 1000 }; // Standard: 1k/hour
    }

    const result = await limiter.checkRateLimit(
      identifier,
      limits.windowMs,
      limits.maxRequests
    );

    // Add headers
    res.set({
      'X-RateLimit-Limit': limits.maxRequests.toString(),
      'X-RateLimit-Remaining': result.remaining.toString(),
      'X-RateLimit-Reset': new Date(result.resetTime).toISOString()
    });

    if (!result.allowed) {
      res.set('Retry-After', result.retryAfter.toString());
      return res.status(429).json({
        success: false,
        message: 'Too many requests',
        retryAfter: result.retryAfter
      });
    }

    next();
  };
}
```

## Documentation et maintenance

### Documentation technique automatisée

```typescript
// OpenAPI/Swagger generation
import { createOpenAPISpec } from '@app/swagger';

// Décorateurs pour auto-documentation
@ApiController('/users')
export class UserController {
  @Get('/')
  @ApiOperation({ summary: 'Get list of users' })
  @ApiQuery({ name: 'page', type: 'number', required: false })
  @ApiQuery({ name: 'limit', type: 'number', required: false })
  @ApiResponse({ status: 200, type: [UserDto] })
  @ApiResponse({ status: 401, description: 'Unauthorized' })
  async getUsers(@Query() query: GetUsersQuery): Promise<UserDto[]> {
    return this.userService.getUsers(query);
  }

  @Post('/')
  @ApiOperation({ summary: 'Create a new user' })
  @ApiBody({ type: CreateUserDto })
  @ApiResponse({ status: 201, type: UserDto })
  @ApiResponse({ status: 400, description: 'Validation error' })
  async createUser(@Body() userData: CreateUserDto): Promise<UserDto> {
    return this.userService.createUser(userData);
  }
}

// Auto-génération de la spec OpenAPI
export function generateOpenAPISpec(): OpenAPIObject {
  return createOpenAPISpec({
    title: 'API Documentation',
    version: '1.0.0',
    description: 'Automatically generated API documentation',
    servers: [
      { url: 'https://api.example.com/v1', description: 'Production' },
      { url: 'https://api-staging.example.com/v1', description: 'Staging' }
    ]
  });
}
```

### Health Checks et monitoring

```typescript
// Health check comprehensive
@Injectable()
export class HealthService {
  constructor(
    private db: Database,
    private redis: Redis,
    private emailService: EmailService
  ) {}

  async getHealthStatus(): Promise<HealthStatus> {
    const checks = await Promise.allSettled([
      this.checkDatabase(),
      this.checkRedis(),
      this.checkEmailService(),
      this.checkExternalAPIs()
    ]);

    const services = {
      database: this.parseHealthCheck(checks[0]),
      redis: this.parseHealthCheck(checks[1]),
      email: this.parseHealthCheck(checks[2]),
      externalAPIs: this.parseHealthCheck(checks[3])
    };

    const overall = Object.values(services).every(s => s.status === 'healthy')
      ? 'healthy'
      : 'unhealthy';

    return {
      status: overall,
      timestamp: new Date().toISOString(),
      version: process.env.APP_VERSION || 'unknown',
      uptime: process.uptime(),
      services,
      system: {
        memory: process.memoryUsage(),
        cpu: process.cpuUsage()
      }
    };
  }

  private async checkDatabase(): Promise<ServiceHealth> {
    try {
      const start = Date.now();
      await this.db.raw('SELECT 1');
      const responseTime = Date.now() - start;

      return {
        status: 'healthy',
        responseTime,
        details: { driver: 'postgresql' }
      };
    } catch (error) {
      return {
        status: 'unhealthy',
        error: error.message,
        details: { driver: 'postgresql' }
      };
    }
  }

  private async checkRedis(): Promise<ServiceHealth> {
    try {
      const start = Date.now();
      await this.redis.ping();
      const responseTime = Date.now() - start;

      return {
        status: 'healthy',
        responseTime,
        details: { version: await this.redis.info('server') }
      };
    } catch (error) {
      return {
        status: 'unhealthy',
        error: error.message
      };
    }
  }
}

// Endpoints de health check
@Controller('/health')
export class HealthController {
  constructor(private healthService: HealthService) {}

  @Get('/')
  async health(): Promise<HealthStatus> {
    return this.healthService.getHealthStatus();
  }

  @Get('/ready')
  async readiness(): Promise<ReadinessStatus> {
    const health = await this.healthService.getHealthStatus();
    const isReady = health.status === 'healthy';

    return {
      ready: isReady,
      checks: health.services
    };
  }

  @Get('/live')
  async liveness(): Promise<LivenessStatus> {
    // Simple check pour Kubernetes liveness probe
    return {
      alive: true,
      timestamp: new Date().toISOString()
    };
  }
}
```

## Migration et évolution

### Stratégie de migration de données

```typescript
// Migration framework
export abstract class Migration {
  abstract readonly version: string;
  abstract readonly description: string;

  abstract up(db: Database): Promise<void>;
  abstract down(db: Database): Promise<void>;
}

// Example migration
export class AddUserPreferencesMigration extends Migration {
  readonly version = '20240315_001';
  readonly description = 'Add user preferences table';

  async up(db: Database): Promise<void> {
    await db.schema.createTable('user_preferences', (table) => {
      table.uuid('id').primary().defaultTo(db.raw('gen_random_uuid()'));
      table.uuid('user_id').references('id').inTable('users').onDelete('CASCADE');
      table.jsonb('preferences').notNullable().defaultTo('{}');
      table.timestamps(true, true);

      table.unique(['user_id']);
      table.index(['user_id']);
    });

    // Populate default preferences for existing users
    await db.raw(`
      INSERT INTO user_preferences (user_id, preferences)
      SELECT id, '{"newsletter": true, "notifications": true}'::jsonb
      FROM users
      WHERE id NOT IN (SELECT user_id FROM user_preferences)
    `);
  }

  async down(db: Database): Promise<void> {
    await db.schema.dropTableIfExists('user_preferences');
  }
}

// Migration runner
export class MigrationRunner {
  constructor(private db: Database) {}

  async runMigrations(): Promise<void> {
    await this.ensureMigrationTable();

    const pendingMigrations = await this.getPendingMigrations();

    for (const migration of pendingMigrations) {
      await this.runMigration(migration);
    }
  }

  private async runMigration(migration: Migration): Promise<void> {
    const trx = await this.db.transaction();

    try {
      await migration.up(trx);
      await trx('migrations').insert({
        version: migration.version,
        description: migration.description,
        executed_at: new Date()
      });

      await trx.commit();
      console.log(`✅ Migration ${migration.version} completed`);
    } catch (error) {
      await trx.rollback();
      console.error(`❌ Migration ${migration.version} failed:`, error);
      throw error;
    }
  }
}
```

### Versioning API

```typescript
// API versioning strategy
export class APIVersioning {
  // Header-based versioning
  @Get('/users')
  @Version(['1', '2'])
  async getUsers(
    @Query() query: GetUsersQuery,
    @Headers('api-version') version?: string
  ): Promise<any> {
    switch (version || '2') {
      case '1':
        return this.getUsersV1(query);
      case '2':
        return this.getUsersV2(query);
      default:
        throw new BadRequestError('Unsupported API version');
    }
  }

  private async getUsersV1(query: GetUsersQuery): Promise<UserV1[]> {
    const users = await this.userService.getUsers(query);
    return users.map(user => this.transformToV1(user));
  }

  private async getUsersV2(query: GetUsersQuery): Promise<UserV2[]> {
    return this.userService.getUsers(query);
  }

  // Transformation pour compatibilité
  private transformToV1(user: User): UserV1 {
    return {
      id: user.id,
      email: user.email,
      name: `${user.firstName} ${user.lastName}`, // V1 avait un champ 'name'
      created: user.createdAt.toISOString()
    };
  }
}

// Deprecation warnings
export function deprecated(version: string, message?: string) {
  return function (target: any, propertyName: string, descriptor: PropertyDescriptor) {
    const method = descriptor.value;

    descriptor.value = function (...args: any[]) {
      console.warn(`⚠️ DEPRECATED: ${target.constructor.name}.${propertyName} is deprecated as of v${version}. ${message || ''}`);

      // Add deprecation header
      if (args[1] && args[1].set) { // Express response object
        args[1].set('X-API-Deprecated', `true; version=${version}`);
      }

      return method.apply(this, args);
    };
  };
}
```

## Conclusion

Cette architecture a été conçue pour répondre aux exigences suivantes :

### Objectifs atteints

- **Scalabilité horizontale** : Architecture microservices avec load balancing
- **Haute disponibilité** : Réplication des services critiques et circuit breakers
- **Performance optimisée** : Cache multi-niveau et optimisations de requêtes
- **Sécurité robuste** : Authentification JWT, validation d'entrées, rate limiting
- **Maintenabilité élevée** : Code modulaire, tests automatisés, documentation

### Métriques de performance cibles

| Métrique | Objectif | Mesure actuelle |
|----------|----------|-----------------|
| **Temps de réponse API** | < 200ms (95e percentile) | [À mesurer] |
| **Disponibilité** | > 99.5% | [À mesurer] |
| **Throughput** | > 1000 req/sec | [À mesurer] |
| **Temps de build** | < 5 minutes | [À mesurer] |
| **Couverture de tests** | > 85% | [À mesurer] |

### Points d'amélioration futurs

1. **Microservices** : Migration vers une architecture plus granulaire
2. **Event Sourcing** : Implémentation pour l'audit et la traçabilité
3. **GraphQL** : API plus flexible pour les clients mobiles
4. **Machine Learning** : Intégration de modèles pour la personnalisation
5. **Observabilité** : Tracing distribué avec OpenTelemetry

### Contacts et responsabilités

- **Architecte principal** : [Nom] - [email]
- **Tech Lead Backend** : [Nom] - [email]
- **Tech Lead Frontend** : [Nom] - [email]
- **DevOps Lead** : [Nom] - [email]

---

**Dernière mise à jour** : [Date]
**Version de l'architecture** : [Version]
**Prochaine revue** : [Date prévue]
**Approuvé par** : [Nom de l'architecte/CTO]# Architecture du système

## Vue d'ensemble

Cette documentation décrit l'architecture technique du projet [nom du projet], ses composants principaux, leurs interactions, et les décisions architecturales prises pour répondre aux exigences fonctionnelles et non-fonctionnelles.

### Objectifs architecturaux

- **Scalabilité** : Capacité à gérer une charge croissante d'utilisateurs
- **Maintenabilité** : Code structuré et facilement modifiable
- **Performance** : Temps de réponse optimaux pour une bonne UX
- **Sécurité** : Protection des données et accès sécurisés
- **Disponibilité** : Haute disponibilité avec un minimum de downtime
- **Évolutivité** : Facilité d'ajout de nouvelles fonctionnalités

## Architecture générale

### Diagramme de haut niveau

```
┌─────────────────────────────────────────────────────┐
│                    UTILISATEURS                     │
│              (Web, Mobile, API)                     │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                LOAD BALANCER                        │
│              (Nginx/HAProxy)                        │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                   CDN/CACHE                         │
│              (CloudFlare/Redis)                     │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                 FRONTEND LAYER                      │
│          (React/Vue/Angular + Static Assets)        │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                 API GATEWAY                         │
│            (Authentication, Routing)                │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│               BACKEND SERVICES                      │
│         ┌─────────┬─────────┬─────────┐            │
│         │   API   │ Worker  │ Cron    │            │
│         │ Service │ Service │ Service │            │
│         └─────────┴─────────┴─────────┘            │
└─────────────────────┬───────────────────────────────┘
                      │
┌─────────────────────▼───────────────────────────────┐
│                DATA LAYER                           │
│    ┌────────────┬────────────┬────────────┐        │
│    │  Database  │   Cache    │  File      │        │
│    │ (Primary)  │  (Redis)   │ Storage    │        │
│    └────────────┴────────────┴────────────┘        │
└─────────────────────────────────────────────────────┘
```

### Patterns architecturaux utilisés

- **Architecture en couches** (Layered Architecture)
- **Microservices** ou **Service-Oriented Architecture (SOA)**
- **Model-View-Controller (MVC)**
- **Repository Pattern** pour l'accès aux données
- **Observer Pattern** pour les événements
- **Circuit Breaker** pour la résilience
- **CQRS** (Command Query Responsibility Segregation) [si applicable]

## Architecture frontend

### Stack technologique

- **Framework** : [React/Vue.js/Angular] v[version]
- **State Management** : [Redux/Vuex/NgRx/Zustand]
- **Router** : [React Router/Vue Router/Angular Router]
- **Styling** : [Tailwind CSS/Styled Components/SCSS]
- **Build Tool** : [Vite/Webpack/Parcel]
- **Testing** : [Jest/Vitest + Testing Library]

### Structure des composants

```
src/
├── components/           # Composants réutilisables
│   ├── ui/              # Composants UI de base (Button, Input)
│   ├── forms/           # Composants de formulaires
│   ├── layout/          # Composants de mise en page
│   └── business/        # Composants métier spécifiques
├── pages/               # Pages de l'application
│   ├── auth/           # Pages d'authentification
│   ├── dashboard/      # Pages du tableau de bord
│   └── settings/       # Pages de paramètres
├── hooks/               # Custom hooks React
├── services/            # Services API et logique métier
├── stores/              # Gestion d'état (Redux/Zustand)
├── utils/               # Fonctions utilitaires
├── types/               # Définitions TypeScript
└── assets/              # Assets statiques
```

### Architecture des composants

```typescript
// Exemple d'architecture de composant
interface ComponentProps {
  // Props typées
}

interface ComponentState {
  // State local
}

// Composant fonctionnel avec hooks
const MyComponent: React.FC<ComponentProps> = ({ prop1, prop2 }) => {
  // Hooks pour l'état local
  const [state, setState] = useState<ComponentState>();

  // Hooks pour l'état global
  const globalState = useSelector(selectSomeData);
  const dispatch = useDispatch();

  // Hooks personnalisés
  const { data, loading, error } = useApiData();

  // Effects
  useEffect(() => {
    // Side effects
  }, []);

  // Event handlers
  const handleAction = useCallback(() => {
    // Action logic
  }, []);

  // Render
  return (
    <div>
      {/* JSX */}
    </div>
