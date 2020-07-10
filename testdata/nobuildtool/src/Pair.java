import java.util.Objects;
import java.util.Optional;
import java.util.function.*;

/**
 * The instance from this class holds two different instance values.
 *
 * to get the value of this instance, use {@link #unify <code>unify</code>} method.
 * For example,
 * <code>
 *     Pair<L, R> pair = Pair.of(...);
 *     L left = pair.unify((l, r) -> l);
 * </code>
 * https://gist.github.com/tamada/930bff6d3a72da2533eacf3a833b7563
 */
public class Pair<L, R> {
    private L left;
    private R right;

    public static <L, R> Pair<L, R> of(L left, R right) {
        return new Pair<>(left, right);
    }

    public Pair<R, L> swap() {
        return of(right, left);
    }

    public void accept(BiConsumer<L, R> action) {
        action.accept(left, right);
    }

    public void accept(Consumer<L> leftAction, Consumer<R> rightAction) {
        leftAction.accept(left);
        rightAction.accept(right);
    }

    public <K> K unify(BiFunction<L, R, K> mapper) {
        return mapper.apply(left, right);
    }

    public <LL, RR> Pair<LL, RR> map(Function<L, LL> leftMapper, Function<R, RR> rightMapper) {
        return of(leftMapper.apply(left),
                rightMapper.apply(right));
    }

    public boolean test(BiPredicate<L, R> predicate) {
        return predicate.test(left, right);
    }

    public Pair<Boolean, Boolean> test(Predicate<L> leftPredicate, Predicate<R> rightPredicate) {
        return map(l -> leftPredicate.test(l),
                r -> rightPredicate.test(r));
    }

    @Override
    public int hashCode() {
        return Objects.hash(left, right);
    }

    @Override
    public boolean equals(Object other) {
        return other instanceof Pair &&
                ((Pair)other).test((l, r) -> Objects.equals(l, left) && Objects.equals(r, right));
    }

    @Override
    public String toString() {
        return String.format("(%s, %s)", left, right);
    }

    private Pair(L left, R right) {
        this.left = Optional.of(left)
                .orElseThrow(() -> new NullPointerException());
        this.right = Optional.of(right)
                .orElseThrow(() -> new NullPointerException());
    }
}
