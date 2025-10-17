namespace Io.Pdl.Infrastructure.Data.RowSet
{
    public class RowSet<TRow>
        where TRow : class, new()
    {
        public virtual TRow NewRecord()
        {
            return new TRow();
        }
    }
}
